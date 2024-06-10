package supplier

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	protoproduct "github.com/fatalistix/trade-organization-proto/gen/go/product"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
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

func (r *Repository) CreateSupplierContext(
	ctx context.Context,
	name string,
) (int32, error) {
	const op = "repository.supplier.CreateSupplier"

	var id int32
	values := map[string]interface{}{
		"name": name,
	}

	err := r.db.NewInsert().
		Model(&values).
		TableExpr("supplier").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new supplier: %w", op, err)
	}

	return id, nil
}

func (r *Repository) SuppliersContext(ctx context.Context) ([]*proto.Supplier, error) {
	const op = "repository.supplier.Suppliers"

	suppliers := make([]supplier, 0)
	err := r.db.NewSelect().
		Column("id", "name", "type").
		TableExpr("supplier").
		Scan(ctx, &suppliers)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to select suppliers: %w", op, err)
	}

	protoSuppliers := make([]*proto.Supplier, 0, len(suppliers))
	for _, supplier := range suppliers {
		protoSupplier, err := supplier.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert supplier: %w", op, err)
		}
		protoSuppliers = append(protoSuppliers, protoSupplier)
	}

	return protoSuppliers, nil
}

func (r *Repository) SupplierContext(ctx context.Context, id int32) (*proto.Supplier, error) {
	const op = "repository.supplier.Supplier"

	var supplier supplier
	err := r.db.NewSelect().
		Column("id", "name", "type").
		TableExpr("supplier").
		Where("id = ?", id).
		Scan(ctx, &supplier)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to select supplier: %w", op, err)
	}

	productSuppliers, err := r.ProductSuppliersContext(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product suppliers: %w", op, err)
	}

	return supplier.ToProtoWith(productSuppliers)
}

func (r *Repository) ProductSuppliersContext(ctx context.Context, supplierID *int32) ([]*proto.ProductSupplier, error) {
	const op = "repository.supplier.ProductSuppliers"

	query := r.db.NewSelect().
		Column("product_id", "supplier_id", "price").
		TableExpr("product_supplier")

	if supplierID != nil {
		query = query.Where("supplier_id = ?", *supplierID)
	}

	productSuppliers := make([]productSupplier, 0)
	err := query.Scan(ctx, &productSuppliers)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get product suppliers: %w", op, err)
	}

	ids := make([]int32, 0, len(productSuppliers))
	for _, productSupplier := range productSuppliers {
		ids = append(ids, productSupplier.ProductID)
	}

	protoProducts := make([]*proto.ProductSupplier, 0, len(productSuppliers))
	for _, productSupplier := range productSuppliers {
		protoProduct, err := productSupplier.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert product supplier to proto: %w", op, err)
		}

		protoProducts = append(protoProducts, protoProduct)
	}

	return protoProducts, nil
}

func (r *Repository) UpdateSupplierContext(ctx context.Context, supplier *proto.Supplier) error {
	const op = "repository.supplier.UpdateSupplier"

	supplierType, err := ProtoSupplierTypeToString(supplier.Type)
	if err != nil {
		return fmt.Errorf("%s: unable to convert supplier type: %w", op, err)
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("%s: unable to start transaction: %w", op, err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	values := map[string]interface{}{
		"name": supplier.Name,
		"type": supplierType,
	}
	_, err = tx.NewUpdate().
		Model(&values).
		TableExpr("supplier").
		Where("id = ?", supplier.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update supplier: %w", op, err)
	}

	_, err = tx.NewDelete().
		TableExpr("product_supplier").
		Where("supplier_id = ?", supplier.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to delete product suppliers: %w", op, err)
	}

	for _, productSupplier := range supplier.Products {
		values := map[string]interface{}{
			"supplier_id": supplier.Id,
			"product_id":  productSupplier.ProductId,
			"price":       core.ProtoMoneyToString(productSupplier.Price),
		}

		_, err = tx.NewInsert().
			Model(&values).
			TableExpr("product_supplier").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("%s: unable to insert product supplier: %w", op, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: unable to commit transaction: %w", op, err)
	}

	return nil
}
