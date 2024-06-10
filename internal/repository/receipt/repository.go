package receipt

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/receipt"
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

func (r *Repository) CreateReceiptWithAccountingContext(
	ctx context.Context,
	sellerId int32,
	products []*proto.ProductReceipt,
	customerId int32,
	receiptingPointWithAccountingId int32,
	receiptingPointWithAccountingType prototradingpoint.ReceiptingPointWithAccountingType,
) (int32, error) {
	const op = "repository.receipt.CreateReceiptWithAccounting"

	return 0, nil
}
