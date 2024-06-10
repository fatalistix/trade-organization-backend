package cancel

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.CancelRequest) (*proto.CancelResponse, error)

type OrderCanceler interface {
	CancelOrderContext(ctx context.Context, id int32) error
}

func MakeCompleteHandlerFunc(log *slog.Logger, completer OrderCanceler) HandlerFunc {
	const op = "grpc.order.handler.complete.MakeCancelHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CancelRequest) (*proto.CancelResponse, error) {
		log.Info("request encoded")

		err := completer.CancelOrderContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to complete order", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to complete order: %s", err)
		}

		log.Info("order completed")

		return &proto.CancelResponse{
			Id: req.Id,
		}, nil
	}
}
