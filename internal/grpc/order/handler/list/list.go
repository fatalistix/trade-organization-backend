package list

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type OrderProvider interface {
	OrdersContext(
		ctx context.Context,
	) ([]*proto.Order, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider OrderProvider) HandlerFunc {
	const op = "grpc.order.handler.list.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		orders, err := provider.OrdersContext(ctx)
		if err != nil {
			log.Error("unable to get orders", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get orders: %s", err)
		}

		log.Debug("orders received", slog.Any("orders", orders))

		return &proto.ListResponse{
			Orders: orders,
		}, nil
	}
}
