package customer

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	"github.com/fatalistix/trade-organization-backend/internal/repository/core"
	protocore "github.com/fatalistix/trade-organization-proto/gen/go/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
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

func (r *Repository) CreateCustomerContext(
	ctx context.Context,
	firstName string,
	lastName string,
	birthDate *protocore.Date,
	phoneNumber string,
) (int32, error) {
	const op = "repository.customer.RegisterCustomer"

	var id int32
	values := map[string]interface{}{
		"first_name":   firstName,
		"last_name":    lastName,
		"birth_date":   core.ProtoDateToString(birthDate),
		"phone_number": phoneNumber,
	}
	err := r.db.NewInsert().
		Model(&values).
		TableExpr("customer").
		Returning("id").
		Scan(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: unable to insert new customer: %w", op, err)
	}

	return id, nil
}

func (r *Repository) CustomerContext(ctx context.Context, id int32) (*proto.Customer, error) {
	const op = "repository.customer.GetCustomer"

	var customer customer
	err := r.db.NewSelect().
		Column("id", "first_name", "last_name", "birth_date", "phone_number").
		TableExpr("customer").
		Where("id = ?", id).
		Scan(ctx, &customer)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get customer: %w", op, err)
	}

	return customer.ToProto()
}

func (r *Repository) CustomersContext(ctx context.Context) ([]*proto.Customer, error) {
	const op = "repository.customer.GetCustomers"

	var customers []customer
	err := r.db.NewSelect().
		Column("id", "first_name", "last_name", "birth_date", "phone_number").
		TableExpr("customer").
		Scan(ctx, &customers)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get customers: %w", op, err)
	}

	protoCustomers := make([]*proto.Customer, 0, len(customers))
	for _, customer := range customers {
		protoCustomer, err := customer.ToProto()
		if err != nil {
			return nil, fmt.Errorf("%s: unable to convert customer: %w", op, err)
		}
		protoCustomers = append(protoCustomers, protoCustomer)
	}

	return protoCustomers, nil
}

func (r *Repository) UpdateCustomerContext(ctx context.Context, customer *proto.Customer) error {
	const op = "repository.customer.UpdateCustomer"

	values := map[string]interface{}{
		"first_name":   customer.FirstName,
		"last_name":    customer.LastName,
		"birth_date":   core.ProtoDateToString(customer.BirthDate),
		"phone_number": customer.PhoneNumber,
	}

	_, err := r.db.NewUpdate().
		Model(&values).
		TableExpr("customer").
		Where("id = ?", customer.Id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: unable to update customer: %w", op, err)
	}

	return nil
}

func (r *Repository) DeleteCustomerContext(ctx context.Context, id int32) error {
	const op = "repository.customer.DeleteCustomer"

	return nil
}
