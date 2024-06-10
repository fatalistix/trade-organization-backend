package createwithaccounting

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/receipt"
	prototradingpoint "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.CreateWithAccountingRequest) (*proto.CreateWithAccountingResponse, error)

type ReceiptCreator interface {
	CreateReceiptWithAccountingContext(
		ctx context.Context,
		sellerId int32,
		products []*proto.ProductReceipt,
		customerId int32,
		receiptingPointWithAccountingId int32,
		receiptingPointWithAccountingType prototradingpoint.ReceiptingPointWithAccountingType,
	) (int32, error)
}

func MakeCreateWithAccountingHandlerFunc(log *slog.Logger, creator ReceiptCreator) HandlerFunc {
	const op = "grpc.receipt.handler.createwithaccounting.MakeCreateWithAccountingHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateWithAccountingRequest) (*proto.CreateWithAccountingResponse, error) {
		log.Debug("request received")

		id, err := creator.CreateReceiptWithAccountingContext(
			ctx,
			req.SellerId,
			req.Products,
			req.CustomerId,
			req.ReceiptingPointWithAccountingId,
			req.ReceiptingPointWithAccountingType,
		)
		if err != nil {
			return nil, err
		}

		return &proto.CreateWithAccountingResponse{
			Id: id,
		}, nil
	}
}
