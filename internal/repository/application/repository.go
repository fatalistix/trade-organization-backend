package application

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	prototradingpoint "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(database *postgres.Database) *Repository {
	return &Repository{
		db: database.DB(),
	}
}

func (r *Repository) CreateApplicationContext(
	ctx context.Context,
	tradingPointID int32,
	tradingPointType prototradingpoint.TradingPointType,
	products []*proto.ProductApplication,
) (int32, error) {
	const op = "repository.application.CreateApplication"

	ttp, err := tradingpoint.ProtoTradingPointTypeToString(tradingPointType)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	var id int32
	values := map[string]interface{}{
		"trading_point_id":   tradingPointID,
		"trading_point_type": ttp,
	}

	err = tx.NewInsert().
		Model(&values).
		TableExpr("application").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new application: %w", op, err)
	}

	for _, product := range products {
		values := map[string]interface{}{
			"application_id": id,
			"product_id":     product.ProductId,
			"quantity":       product.Quantity,
		}
		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("product_application").
			Exec(ctx)
		if err != nil {
			return 0, fmt.Errorf("%s: unable to insert new application product: %w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return id, nil
}

func (r *Repository) ApplicationContext(ctx context.Context, id int32) (*proto.Application, error) {
	const op = "repository.application.Application"

	var application application
	err := r.db.NewSelect().
		Column("id", "trading_point_id", "trading_point_type", "created_at").
		TableExpr("application").
		Where("id = ?", id).
		Scan(ctx, &application)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get application: %w", op, err)
	}

	productApplications, err := r.ProductApplicationsContext(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product applications: %w", op, err)
	}

	return application.ToProtoWith(productApplications)
}

func (r *Repository) ProductApplicationsContext(ctx context.Context, applicationID *int32) ([]*proto.ProductApplication, error) {
	const op = "repository.application.ProductApplications"

	query := r.db.NewSelect().
		Column("product_id", "quantity").
		TableExpr("product_application")

	if applicationID != nil {
		query = query.Where("application_id = ?", *applicationID)
	}

	productApplications := make([]*productApplication, 0)
	err := query.Scan(ctx, &productApplications)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product applications: %w", op, err)
	}

	protoProductApplications := make([]*proto.ProductApplication, 0, len(productApplications))
	for _, productApplication := range productApplications {
		protoProductApplication := productApplication.ToProto()
		protoProductApplications = append(protoProductApplications, protoProductApplication)
	}

	return protoProductApplications, nil
}

func (r *Repository) ApplicationsContext(ctx context.Context) ([]*proto.Application, error) {
	const op = "repository.application.Applications"

	applications := make([]application, 0)
	err := r.db.NewSelect().
		Column("id", "trading_point_id", "trading_point_type", "created_at").
		TableExpr("application").
		Scan(ctx, &applications)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to select applications: %w", op, err)
	}

	protoApplications := make([]*proto.Application, 0, len(applications))
	for _, application := range applications {
		protoApplication, err := application.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert application: %w", op, err)
		}
		protoApplications = append(protoApplications, protoApplication)
	}

	return protoApplications, nil
}

func (r *Repository) UpdateApplicationContext(ctx context.Context, application *proto.Application) error {
	const op = "repository.application.UpdateApplication"

	tradingPointType, err := tradingpoint.ProtoTradingPointTypeToString(application.TradingPointType)
	if err != nil {
		return fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	values := map[string]interface{}{
		"trading_point_id":   application.TradingPointId,
		"trading_point_type": tradingPointType,
		"order_id":           application.OrderId,
	}
	_, err = tx.NewUpdate().
		Model(&values).
		TableExpr("application").
		Where("id = ?", application.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update application: %w", op, err)
	}

	_, err = tx.NewDelete().
		TableExpr("product_application").
		Where("application_id = ?", application.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to delete old application products: %w", op, err)
	}

	for _, productApplication := range application.Products {
		values := map[string]interface{}{
			"application_id": application.Id,
			"product_id":     productApplication.ProductId,
			"quantity":       productApplication.Quantity,
		}

		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("product_application").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("%s: unable to insert product application: %w", op, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return nil
}
