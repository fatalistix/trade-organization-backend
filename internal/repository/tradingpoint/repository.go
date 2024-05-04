package tradingpoint

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/model/tradingpoint"
	"github.com/uptrace/bun"

	model "github.com/fatalistix/trade-organization-backend/internal/model/tradingpoint"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(database *postgres.Database) *Repository {
	return &Repository{db: database.DB()}
}

func (r *Repository) RegisterNewTradingPoint(
	ctx context.Context,
	t tradingpoint.Type,
	areaPlot float64,
	rentalCharge *core.Money,
	counterCount int32,
	address string,
) (int32, error) {
	const op = "repository.tradingpoint.RegisterNewTradingPoint"

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
	case tradingpoint.TypeDepartmentStore:
		err = r.registerDepartmentStore(ctx, tx, id)
	case tradingpoint.TypeStore:
		err = r.registerStore(ctx, tx, id)
	case tradingpoint.TypeKiosk:
		err = r.registerKiosk(ctx, tx, id)
	case tradingpoint.TypeTray:
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

	if err := r.registerHallContainer(ctx, tx, id, "store"); err != nil {
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

	if err := r.registerPlaceOfWork(ctx, tx, id, "kiosk"); err != nil {
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

	if err := r.registerPlaceOfWork(ctx, tx, id, "tray"); err != nil {
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

func (r *Repository) registerHallContainer(ctx context.Context, tx bun.Tx, id int32, t string) error {
	const op = "repository.tradingpoint.registerHallContainer"

	values := map[string]interface{}{
		"id":   id,
		"type": t,
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("hall_container").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new hall container: %w", op, err)
	}

	return nil
}

func (r *Repository) registerPlaceOfWork(ctx context.Context, tx bun.Tx, id int32, t string) error {
	const op = "repository.tradingpoint.registerPlaceOfWork"

	values := map[string]interface{}{
		"id":   id,
		"type": t,
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("place_of_work").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to insert new place of work: %w", op, err)
	}

	return nil
}

func (r *Repository) List(ctx context.Context) ([]*model.TradingPoint, error) {
	const op = "repository.tradingpoint.List"

	tradingPoints := make([]*tradingPoint, 0)

	err := r.db.NewSelect().
		Column("id", "type", "area_plot", "rental_charge", "counter_count", "address").
		TableExpr("trading_point").
		Scan(ctx, &tradingPoints)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to list trading points: %w", op, err)
	}

	mtps := make([]*model.TradingPoint, 0, len(tradingPoints))
	for _, tp := range tradingPoints {
		mtp, err := tp.ToModel()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert trading point to model: %w", op, err)
		}
		mtps = append(mtps, mtp)
	}

	return mtps, nil
}
