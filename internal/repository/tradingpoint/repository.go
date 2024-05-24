package tradingpoint

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/department_store"
	modelhall "github.com/fatalistix/trade-organization-backend/internal/domain/model/hall"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/receipting_point_with_accounting"
	modelsection "github.com/fatalistix/trade-organization-backend/internal/domain/model/section"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(database *postgres.Database) *Repository {
	return &Repository{db: database.DB()}
}

func (r *Repository) RegisterTradingPoint(
	ctx context.Context,
	t trading_point.Type,
	areaPlot float64,
	rentalCharge core.Money,
	counterCount int32,
	address string,
) (int32, error) {
	const op = "repository.tradingpoint.RegisterTradingPoint"

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf("%s: unable to start tx: %w", op, err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	var id int32
	values := map[string]interface{}{
		"type":          t,
		"area_plot":     areaPlot,
		"rental_charge": rentalCharge.String(),
		"counter_count": counterCount,
		"address":       address,
	}
	err = tx.NewInsert().
		Model(&values).
		TableExpr("trading_point").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new trading point: %w", op, err)
	}

	switch t {
	case trading_point.TypeDepartmentStore:
		err = r.registerDepartmentStore(ctx, tx, id)
	case trading_point.TypeStore:
		err = r.registerStore(ctx, tx, id)
	case trading_point.TypeKiosk:
		err = r.registerKiosk(ctx, tx, id)
	case trading_point.TypeTray:
		err = r.registerTray(ctx, tx, id)
	default:
		err = fmt.Errorf("%s: unknown type %s", op, t)
	}

	if err != nil {
		return 0, fmt.Errorf("%s: unable to register new trading point: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%s: unable to commit tx: %w", op, err)
	}

	return id, nil
}

func (r *Repository) registerDepartmentStore(ctx context.Context, tx bun.Tx, id int32) error {
	const op = "repository.tradingpoint.registerDepartmentStore"

	if err := r.registerReceiptingPointWithAccounting(ctx, tx, id, "department_store"); err != nil {
		return fmt.Errorf("%s: unable to register department store: %w", op, err)
	}

	values := map[string]interface{}{
		"id":   id,
		"type": "department_store",
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("department_store").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new department store: %w", op, err)
	}

	return nil
}

func (r *Repository) registerStore(ctx context.Context, tx bun.Tx, id int32) error {
	const op = "repository.tradingpoint.registerStore"

	if err := r.registerReceiptingPointWithAccounting(ctx, tx, id, "store"); err != nil {
		return fmt.Errorf("%s: unable to register store: %w", op, err)
	}

	if _, err := r.registerHallContainer(ctx, tx, &id, "store"); err != nil {
		return fmt.Errorf("%s: unable to register store: %w", op, err)
	}

	values := map[string]interface{}{
		"id":   id,
		"type": "store",
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("store").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new store: %w", op, err)
	}

	return nil
}

func (r *Repository) registerKiosk(ctx context.Context, tx bun.Tx, id int32) error {
	const op = "repository.tradingpoint.registerKiosk"

	if err := r.registerReceiptingPointWithoutAccounting(ctx, tx, id, "kiosk"); err != nil {
		return fmt.Errorf("%s: unable to register kiosk: %w", op, err)
	}

	if _, err := r.registerPlaceOfWork(ctx, tx, &id, "kiosk"); err != nil {
		return fmt.Errorf("%s: unable to register kiosk: %w", op, err)
	}

	values := map[string]interface{}{
		"id":   id,
		"type": "kiosk",
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("kiosk").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new kiosk: %w", op, err)
	}

	return nil
}

func (r *Repository) registerTray(ctx context.Context, tx bun.Tx, id int32) error {
	const op = "repository.tradingpoint.registerTray"

	if err := r.registerReceiptingPointWithoutAccounting(ctx, tx, id, "tray"); err != nil {
		return fmt.Errorf("%s: unable to register tray: %w", op, err)
	}

	if _, err := r.registerPlaceOfWork(ctx, tx, &id, "tray"); err != nil {
		return fmt.Errorf("%s: unable to register tray: %w", op, err)
	}

	values := map[string]interface{}{
		"id":   id,
		"type": "tray",
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("tray").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new tray: %w", op, err)
	}

	return nil
}

func (r *Repository) registerReceiptingPointWithAccounting(ctx context.Context, tx bun.Tx, id int32, t string) error {
	const op = "repository.tradingpoint.registerReceiptingPointWithAccounting"

	values := map[string]interface{}{
		"id":   id,
		"type": t,
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("receipting_point_with_accounting").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new receipting point with accounting: %w", op, err)
	}

	return nil
}

func (r *Repository) registerReceiptingPointWithoutAccounting(ctx context.Context, tx bun.Tx, id int32, t string) error {
	const op = "repository.tradingpoint.registerReceiptingPointWithoutAccounting"

	values := map[string]interface{}{
		"id":   id,
		"type": t,
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("receipting_point_without_accounting").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new receipting point without accounting: %w", op, err)
	}

	return nil
}

func (r *Repository) registerHallContainer(ctx context.Context, tx bun.Tx, id *int32, t string) (int32, error) {
	const op = "repository.tradingpoint.registerHallContainer"

	values := map[string]interface{}{
		"type": t,
	}
	if id != nil {
		values["id"] = *id
	}

	var resultID int32
	err := tx.NewInsert().
		Model(&values).
		TableExpr("hall_container").
		Returning("id").
		Scan(ctx, &resultID)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new hall container: %w", op, err)
	}

	return resultID, nil
}

func (r *Repository) registerPlaceOfWork(ctx context.Context, tx bun.Tx, id *int32, t string) (int32, error) {
	const op = "repository.tradingpoint.registerPlaceOfWork"

	values := map[string]interface{}{
		"type": t,
	}
	if id != nil {
		values["id"] = *id
	}

	var resultID int32
	err := tx.NewInsert().
		Model(&values).
		TableExpr("place_of_work").
		Returning("id").
		Scan(ctx, &resultID)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new place of work: %w", op, err)
	}

	return resultID, nil
}

func (r *Repository) TradingPointsContext(ctx context.Context) ([]trading_point.TradingPoint, error) {
	const op = "repository.tradingpoint.List"

	tradingPoints := make([]*tradingPoint, 0)

	err := r.db.NewSelect().
		Column("id", "type", "area_plot", "rental_charge", "counter_count", "address").
		TableExpr("trading_point").
		Scan(ctx, &tradingPoints)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to list trading points: %w", op, err)
	}

	mtps := make([]trading_point.TradingPoint, 0, len(tradingPoints))
	for _, tp := range tradingPoints {
		mtp, err := tp.ToModel()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert trading point to model: %w", op, err)
		}
		mtps = append(mtps, mtp)
	}

	return mtps, nil
}

func (r *Repository) TradingPointContext(ctx context.Context, id int32, t trading_point.Type) (trading_point.TradingPoint, error) {
	const op = "repository.tradingpoint.TradingPoint"

	var tradingPoint tradingPoint
	err := r.db.NewSelect().
		Column("id", "type", "area_plot", "rental_charge", "counter_count", "address").
		TableExpr("trading_point").
		Where("id = ?", id).
		Where("type = ?", t.String()).
		Scan(ctx, &tradingPoint)
	if err != nil {
		return trading_point.TradingPoint{}, fmt.Errorf("%s: unable to get trading point: %w", op, err)
	}

	return tradingPoint.ToModel()
}

func (r *Repository) AddHallContext(
	ctx context.Context,
	hallContainerID int32,
	hallContainerType hall_container.Type,
	tradingPointID int32,
	tradingPointType trading_point.Type,
) (int32, error) {
	const op = "repository.tradingpoint.AddHallContext"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	id, err := r.registerPlaceOfWork(ctx, tx, nil, "hall")
	if err != nil {
		return 0, fmt.Errorf("%s: unable to register place of work: %w", op, err)
	}

	values := map[string]interface{}{
		"id":                  id,
		"type":                "hall",
		"hall_container_id":   hallContainerID,
		"hall_container_type": hallContainerType.String(),
		"trading_point_id":    tradingPointID,
		"trading_point_type":  tradingPointType.String(),
	}
	_, err = tx.NewInsert().
		Model(&values).
		TableExpr("hall").
		Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new hall: %w", op, err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return id, nil
}

func (r *Repository) AddSectionContext(
	ctx context.Context,
	departmentStoreID int32,
) (int32, error) {
	const op = "repository.tradingpoint.AddSectionContext"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	id, err := r.registerHallContainer(ctx, tx, nil, "section")
	if err != nil {
		return 0, fmt.Errorf("%s: unable to register hall container: %w", op, err)
	}

	values := map[string]interface{}{
		"id":                    id,
		"type":                  "section",
		"department_store_id":   departmentStoreID,
		"department_store_type": "department_store",
	}
	_, err = tx.NewInsert().
		Model(&values).
		TableExpr("section").
		Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new section: %w", op, err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return id, nil
}

func (r *Repository) DepartmentStoreContext(ctx context.Context, id int32) (department_store.DepartmentStore, error) {
	const op = "repository.tradingpoint.DepartmentStoreContext"

	tradingPoint, err := r.TradingPointContext(ctx, id, trading_point.TypeDepartmentStore)
	if err != nil {
		return department_store.DepartmentStore{}, fmt.Errorf("%s: unable to get trading point: %w", op, err)
	}

	receiptingPointWithAccounting, err := r.ReceiptingPointWithAccountingContext(ctx, id, receipting_point_with_accounting.TypeDepartmentStore)
	if err != nil {
		return department_store.DepartmentStore{}, fmt.Errorf("%s: unable to get receipting point with accounting: %w", op, err)
	}

	tradingPointType := trading_point.TypeDepartmentStore
	halls, err := r.HallsContext(ctx, &id, &tradingPointType)
	if err != nil {
		return department_store.DepartmentStore{}, fmt.Errorf("%s: unable to get halls: %w", op, err)
	}

	sections, err := r.SectionsContext(ctx, &id)
	if err != nil {
		return department_store.DepartmentStore{}, fmt.Errorf("%s: unable to get sections: %w", op, err)
	}

	return department_store.DepartmentStore{
		ID:                            id,
		TradingPoint:                  tradingPoint,
		ReceiptingPointWithAccounting: receiptingPointWithAccounting,
		Halls:                         halls,
		Sections:                      sections,
	}, nil
}

func (r *Repository) ReceiptingPointWithAccountingContext(ctx context.Context, id int32, t receipting_point_with_accounting.Type) (receipting_point_with_accounting.ReceiptingPointWithAccounting, error) {
	const op = "repository.tradingpoint.ReceiptingPointWithAccounting"

	var receiptingPointWithAccounting receiptingPointWithAccounting
	err := r.db.NewSelect().
		Column("id", "type").
		TableExpr("receipting_point_with_accounting").
		Where("id = ?", id).
		Where("type = ?", t.String()).
		Scan(ctx, &receiptingPointWithAccounting)
	if err != nil {
		return receipting_point_with_accounting.ReceiptingPointWithAccounting{}, fmt.Errorf("%s: unable to get receipting point with accounting: %w", op, err)
	}

	return receiptingPointWithAccounting.ToModel()
}

func (r *Repository) HallsContext(ctx context.Context, tradingPointID *int32, tradingPointType *trading_point.Type) ([]modelhall.Hall, error) {
	const op = "repository.tradingpoint.HallsContext"

	halls := make([]hall, 0)
	err := r.db.NewSelect().
		Column("id", "type", "hall_container_id", "hall_container_type", "trading_point_id", "trading_point_type").
		TableExpr("hall").
		Where("trading_point_id = ?", tradingPointID).
		Where("trading_point_type = ?", tradingPointType.String()).
		Scan(ctx, &halls)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get halls: %w", op, err)
	}

	modelHalls := make([]modelhall.Hall, 0, len(halls))
	for _, h := range halls {
		mh, err := h.ToModel()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert hall: %w", op, err)
		}
		modelHalls = append(modelHalls, mh)
	}

	return modelHalls, nil
}

func (r *Repository) SectionsContext(ctx context.Context, departmentStoreID *int32) ([]modelsection.Section, error) {
	const op = "repository.tradingpoint.SectionsContext"

	sections := make([]section, 0)
	err := r.db.NewSelect().
		Column("id", "type", "department_store_id").
		TableExpr("section").
		Where("department_store_id = ?", departmentStoreID).
		Scan(ctx, &sections)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get sections: %w", op, err)
	}

	modelSections := make([]modelsection.Section, 0, len(sections))
	for _, s := range sections {
		modelSections = append(modelSections, s.ToModel())
	}

	return modelSections, nil
}
