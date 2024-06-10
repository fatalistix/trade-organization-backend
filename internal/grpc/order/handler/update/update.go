package update

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type OrderUpdater interface {
	UpdateOrderContext(ctx context.Context, order *proto.Order) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater OrderUpdater) HandlerFunc {
	const op = "grpc.order.handler.update.MakeUpdateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded")

		err := updater.UpdateOrderContext(ctx, req.Order)
		if err != nil {
			log.Error("unable to update order", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update order: %s", err)
		}

		log.Info("order updated")

		return &proto.UpdateResponse{
			Id: req.Order.Id,
		}, nil
	}
}
