package product

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/product"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(database *postgres.Database) *Repository {
	return &Repository{db: database.DB()}
}

func (r *Repository) CreateProductContext(
	ctx context.Context,
	name string,
	description string,
) (int32, error) {
	const op = "repository.product.CreateProduct"

	var id int32
	values := map[string]interface{}{
		"name":        name,
		"description": description,
	}

	err := r.db.NewInsert().
		Model(&values).
		TableExpr("product").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new product: %w", op, err)
	}

	return id, nil
}

func (r *Repository) ProductsContext(
	ctx context.Context,
	ids *[]int32,
) ([]*proto.Product, error) {
	const op = "repository.product.Products"

	query := r.db.NewSelect().
		Column("id", "name", "description").
		TableExpr("product")

	if ids != nil {
		query = query.Where("id IN (?)", bun.In(*ids))
	}

	products := make([]*product, 0)
	err := query.Scan(ctx, &products)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get products: %w", op, err)
	}

	protoProducts := make([]*proto.Product, 0, len(products))

	for _, product := range products {
		protoProducts = append(protoProducts, product.ToProto())
	}

	return protoProducts, nil
}

func (r *Repository) ProductContext(
	ctx context.Context,
	id int32,
) (*proto.Product, error) {
	const op = "repository.product.Product"

	var product product
	err := r.db.NewSelect().
		Column("id", "name", "description").
		TableExpr("product").
		Where("id = ?", id).
		Scan(ctx, &product)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product: %w", op, err)
	}

	return product.ToProto(), nil
}

func (r *Repository) UpdateProductContext(
	ctx context.Context,
	product *proto.Product,
) error {
	const op = "repository.product.UpdateProduct"

	values := map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
	}
	_, err := r.db.NewUpdate().
		Model(&values).
		TableExpr("product").
		Where("id = ?", product.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update product: %w", op, err)
	}

	return nil
}
