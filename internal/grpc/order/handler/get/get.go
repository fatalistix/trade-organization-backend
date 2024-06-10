package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.OrderRequest) (*proto.OrderResponse, error)

type OrderProvider interface {
	OrderContext(ctx context.Context, id int32) (*proto.Order, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider OrderProvider) HandlerFunc {
	const op = "grpc.order.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
		log.Info("request encoded")

		order, err := provider.OrderContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get order", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get order: %s", err)
		}

		log.Info("order received", slog.Any("order", order))

		return &proto.OrderResponse{
			Order: order,
		}, nil
	}
}
