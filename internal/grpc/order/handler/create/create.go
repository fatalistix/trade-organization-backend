package create

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error)

type OrderCreator interface {
	CreateOrderContext(
		ctx context.Context,
		supplierID int32,
		products []*proto.ProductOrder,
		applicationIds []int32,
	) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator OrderCreator) HandlerFunc {
	const op = "grpc.order.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		log.Info("request encoded")

		id, err := creator.CreateOrderContext(ctx, req.SupplierId, req.Products, req.ApplicationIds)
		if err != nil {
			log.Error("unable to create order", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to create order: %s", err)
		}

		log.Info("order created")

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
