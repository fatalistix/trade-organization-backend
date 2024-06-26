package tradingpoint

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	protoproduct "github.com/fatalistix/trade-organization-proto/gen/go/product"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"github.com/uptrace/bun"
)

type ProductProvider interface {
	ProductsContext(ctx context.Context, ids *[]int32) ([]*protoproduct.Product, error)
}

type Repository struct {
	db              *bun.DB
	productProvider ProductProvider
}

func NewRepository(database *postgres.Database, productProvider ProductProvider) *Repository {
	return &Repository{
		db:              database.DB(),
		productProvider: productProvider,
	}
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

func (r *Repository) TradingPointContext(ctx context.Context, id int32, t proto.TradingPointType) (*proto.TradingPoint, error) {
	const op = "repository.tradingpoint.TradingPoint"

	tradingPointType, err := ProtoTradingPointTypeToString(t)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
	}

	var tradingPoint tradingPoint
	err = r.db.NewSelect().
		Column("id", "type", "area_plot", "rental_charge", "counter_count", "address").
		TableExpr("trading_point").
		Where("id = ?", id).
		Where("type = ?", tradingPointType).
		Scan(ctx, &tradingPoint)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get trading point: %w", op, err)
	}

	productTradingPoints, err := r.ProductTradingPointsContext(ctx, &id, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product trading points: %w", op, err)
	}

	return tradingPoint.ToProtoWith(productTradingPoints)
}

func (r *Repository) ProductTradingPointsContext(ctx context.Context, tradingPointID *int32, tradingPointType *proto.TradingPointType) ([]*proto.ProductTradingPoint, error) {
	const op = "repository.tradingpoint.ProductTradingPoints"

	query := r.db.NewSelect().
		Column("id", "trading_point_id", "trading_point_type", "product_id", "quantity", "price").
		TableExpr("product_trading_point")

	if tradingPointID != nil {
		query = query.Where("trading_point_id = ?", *tradingPointID)
	}

	if tradingPointType != nil {
		tradingPointTypeString, err := ProtoTradingPointTypeToString(*tradingPointType)
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
		}
		query = query.Where("trading_point_type = ?", tradingPointTypeString)
	}

	productTradingPoints := make([]productTradingPoint, 0)
	err := query.Scan(ctx, &productTradingPoints)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product trading points: %w", op, err)
	}

	ids := make([]int32, 0, len(productTradingPoints))
	for _, productTradingPoint := range productTradingPoints {
		ids = append(ids, productTradingPoint.ProductID)
	}

	protoProducts := make([]*proto.ProductTradingPoint, 0, len(productTradingPoints))
	for _, productTradingPoint := range productTradingPoints {
		protoProduct, err := productTradingPoint.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert product trading point to proto: %w", op, err)
		}
		protoProducts = append(protoProducts, protoProduct)
	}

	return protoProducts, nil
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

func (r *Repository) DepartmentStoreContext(ctx context.Context, id int32) (*proto.DepartmentStore, error) {
	const op = "repository.tradingpoint.DepartmentStoreContext"

	tradingPoint, err := r.TradingPointContext(ctx, id, proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get trading point: %w", op, err)
	}

	receiptingPointWithAccounting, err := r.ReceiptingPointWithAccountingContext(
		ctx, id, proto.ReceiptingPointWithAccountingType_RECEIPTING_POINT_WITH_ACCOUNTING_DEPARTMENT_STORE,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get receipting point with accounting: %w", op, err)
	}

	tradingPointType := proto.TradingPointType_TRADING_POINT_DEPARTMENT_STORE
	halls, err := r.HallsContext(ctx, &id, &tradingPointType)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get halls: %w", op, err)
	}

	sections, err := r.SectionsContext(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get sections: %w", op, err)
	}

	return &proto.DepartmentStore{
		Id:                            id,
		TradingPoint:                  tradingPoint,
		ReceiptingPointWithAccounting: receiptingPointWithAccounting,
		Halls:                         halls,
		Sections:                      sections,
	}, nil
}

func (r *Repository) ReceiptingPointWithAccountingContext(ctx context.Context, id int32, t proto.ReceiptingPointWithAccountingType) (*proto.ReceiptingPointWithAccounting, error) {
	const op = "repository.tradingpoint.ReceiptingPointWithAccounting"

	receiptingPointWithAccountingTypeString, err := ProtoReceiptingPointWithAccountingTypeToString(t)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to convert receipting point with accounting type: %w", op, err)
	}

	var receiptingPointWithAccounting receiptingPointWithAccounting
	err = r.db.NewSelect().
		Column("id", "type").
		TableExpr("receipting_point_with_accounting").
		Where("id = ?", id).
		Where("type = ?", receiptingPointWithAccountingTypeString).
		Scan(ctx, &receiptingPointWithAccounting)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get receipting point with accounting: %w", op, err)
	}

	return receiptingPointWithAccounting.ToProto()
}

func (r *Repository) HallsContext(ctx context.Context, tradingPointID *int32, tradingPointType *proto.TradingPointType) ([]*proto.Hall, error) {
	const op = "repository.tradingpoint.HallsContext"

	query := r.db.NewSelect().
		Column("id", "type", "hall_container_id", "hall_container_type", "trading_point_id", "trading_point_type").
		TableExpr("hall")

	if tradingPointID != nil {
		query = query.Where("trading_point_id = ?", *tradingPointID)
	}

	if tradingPointType != nil {
		tradingPointTypeString, err := ProtoTradingPointTypeToString(*tradingPointType)
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
		}
		query = query.Where("trading_point_type = ?", tradingPointTypeString)
	}

	halls := make([]hall, 0)

	err := query.Scan(ctx, &halls)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get halls: %w", op, err)
	}

	protoHalls := make([]*proto.Hall, 0, len(halls))
	for _, h := range halls {
		ph, err := h.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert hall: %w", op, err)
		}
		protoHalls = append(protoHalls, ph)
	}

	return protoHalls, nil
}

func (r *Repository) SectionsContext(ctx context.Context, departmentStoreID *int32) ([]*proto.Section, error) {
	const op = "repository.tradingpoint.SectionsContext"

	query := r.db.NewSelect().
		Column("id", "type", "department_store_id").
		TableExpr("section")

	if departmentStoreID != nil {
		query = query.Where("department_store_id = ?", *departmentStoreID)
	}

	sections := make([]section, 0)

	err := query.Scan(ctx, &sections)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get sections: %w", op, err)
	}

	protoSections := make([]*proto.Section, 0, len(sections))
	for _, s := range sections {
		protoSections = append(protoSections, s.ToProto())
	}

	return protoSections, nil
}
