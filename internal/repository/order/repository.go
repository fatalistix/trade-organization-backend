package order

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	"github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
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

func (r *Repository) CreateOrderContext(
	ctx context.Context,
	supplierID int32,
	products []*proto.ProductOrder,
	applicationIds []int32,
) (int32, error) {
	const op = "repository.order.CreateOrder"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	var id int32
	values := map[string]interface{}{
		"supplier_id": supplierID,
	}

	err = tx.NewInsert().
		Model(&values).
		TableExpr(`"order"`).
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new order: %w", op, err)
	}

	for _, product := range products {
		price := core.ProtoMoneyToString(product.Price)

		values := map[string]interface{}{
			"order_id":   id,
			"product_id": product.ProductId,
			"quantity":   product.Quantity,
			"price":      price,
		}
		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("product_order").
			Exec(ctx)
		if err != nil {
			return 0, fmt.Errorf("%s: unable to insert new order product: %w", op, err)
		}
	}

	for _, applicationID := range applicationIds {
		values := map[string]interface{}{
			"order_id": id,
		}
		_, err = tx.NewUpdate().
			Model(&values).
			TableExpr("application").
			Where("id = ?", applicationID).
			Exec(ctx)
		if err != nil {
			return 0, fmt.Errorf("%s: unable to insert new order application: %w", op, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return id, nil
}

func (r *Repository) OrderContext(ctx context.Context, id int32) (*proto.Order, error) {
	const op = "repository.order.GetOrder"

	var order order
	err := r.db.NewSelect().
		Column("id", "supplier_id", "created_at", "status", "completed_at", "canceled_at").
		TableExpr(`"order"`).
		Where("id = ?", id).
		Scan(ctx, &order)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get order: %w", op, err)
	}

	productOrders, err := r.ProductOrdersContext(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product orders: %w", op, err)
	}

	distributions, err := r.DistributionsContext(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get distributions: %w", op, err)
	}

	return order.ToProtoWith(productOrders, distributions)
}

func (r *Repository) ProductOrdersContext(ctx context.Context, orderID *int32) ([]*proto.ProductOrder, error) {
	const op = "repository.order.GetProductOrders"

	query := r.db.NewSelect().
		Column("product_id", "quantity", "price").
		TableExpr("product_order")

	if orderID != nil {
		query = query.Where("order_id = ?", *orderID)
	}

	productOrders := make([]productOrder, 0)
	err := query.Scan(ctx, &productOrders)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product orders: %w", op, err)
	}

	protoProductOrders := make([]*proto.ProductOrder, 0)
	for _, productOrder := range productOrders {
		protoProductOrder, err := productOrder.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert product order to proto: %w", op, err)
		}
		protoProductOrders = append(protoProductOrders, protoProductOrder)
	}

	return protoProductOrders, nil
}

func (r *Repository) DistributionsContext(ctx context.Context, orderID *int32) ([]*proto.Distribution, error) {
	const op = "repository.order.GetDistributions"

	query := r.db.NewSelect().
		Column("distribution_id", "quantity", "price").
		TableExpr("distribution_order")

	if orderID != nil {
		query = query.Where("order_id = ?", *orderID)
	}

	distributions := make([]distribution, 0)
	err := query.Scan(ctx, &distributions)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get distributions: %w", op, err)
	}

	protoDistributions := make([]*proto.Distribution, 0)
	for _, distributionOrder := range distributions {
		protoDistribution, err := distributionOrder.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert distribution order to proto: %w", op, err)
		}
		protoDistributions = append(protoDistributions, protoDistribution)
	}

	return protoDistributions, nil
}

func (r *Repository) OrdersContext(ctx context.Context) ([]*proto.Order, error) {
	const op = "repository.order.GetOrders"

	orders := make([]order, 0)
	err := r.db.NewSelect().
		Column("id", "supplier_id", "created_at", "status", "completed_at", "canceled_at").
		TableExpr(`"order"`).
		Scan(ctx, &orders)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get orders: %w", op, err)
	}

	protoOrders := make([]*proto.Order, 0)
	for _, order := range orders {
		protoOrder, err := order.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert order to proto: %w", op, err)
		}
		protoOrders = append(protoOrders, protoOrder)
	}

	return protoOrders, nil
}

func (r *Repository) UpdateOrderContext(ctx context.Context, order *proto.Order) error {
	const op = "repository.order.UpdateOrder"

	status, err := ProtoOrderStatusToString(order.Status)
	if err != nil {
		return fmt.Errorf("%s: unable to convert proto order status to string: %w", op, err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	values := map[string]interface{}{
		"supplier_id": order.SupplierId,
		"created_at":  order.CreatedAt.AsTime(),
		"status":      status,
	}
	if order.CompletedAt != nil {
		values["completed_at"] = order.CompletedAt.AsTime()
	}
	if order.CanceledAt != nil {
		values["canceled_at"] = order.CanceledAt.AsTime()
	}

	_, err = tx.NewUpdate().
		Model(&values).
		TableExpr(`"order"`).
		Where("id = ?", order.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update order: %w", op, err)
	}

	_, err = tx.NewDelete().
		TableExpr("product_order").
		Where("order_id = ?", order.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to delete product orders: %w", op, err)
	}

	for _, productOrder := range order.Products {
		values := map[string]interface{}{
			"order_id":   order.Id,
			"product_id": productOrder.ProductId,
			"quantity":   productOrder.Quantity,
			"price":      core.ProtoMoneyToString(productOrder.Price),
		}

		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("product_order").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("%s: unable to insert product order: %w", op, err)
		}
	}

	for _, distribution := range order.Distributions {
		tradingPointType, err := tradingpoint.ProtoTradingPointTypeToString(distribution.TradingPointType)
		if err != nil {
			return fmt.Errorf("%s: unable to convert trading point type: %w", op, err)
		}

		values := map[string]interface{}{
			"order_id":           order.Id,
			"product_id":         distribution.ProductId,
			"quantity":           distribution.Quantity,
			"trading_point_id":   distribution.TradingPointId,
			"trading_point_type": tradingPointType,
		}

		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("distribution_order").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("%s: unable to insert distribution order: %w", op, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return nil
}

func (r *Repository) CompleteOrderContext(ctx context.Context, id int32) error {
	const op = "repository.order.CompleteOrder"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.NewUpdate().
		TableExpr(`"order"`).
		Set("status = 'completed'").
		Set("completed_at = now()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update order: %w", op, err)
	}

	productOrders := make([]productOrder, 0)
	err = tx.NewSelect().
		Column("product_id", "quantity").
		Table("product_order").
		Where("order_id = ?", id).
		Scan(ctx, &productOrders)
	if err != nil {
		return fmt.Errorf("%s: unable to get product orders: %w", op, err)
	}

	for _, productOrder := range productOrders {
		_, err = tx.NewUpdate().
			Table("product_trading_point").
			Set("quantity = quantity + ?", productOrder.Quantity).
			Where("product_id = ?", productOrder.ProductID).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("%s: unable to update product stock: %w", op, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return nil
}

func (r *Repository) CancelOrderContext(ctx context.Context, id int32) error {
	const op = "repository.order.CancelOrder"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.NewUpdate().
		TableExpr(`"order"`).
		Set("status = 'canceled'").
		Set("canceled_at = now()").
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update order: %w", op, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return nil
}
