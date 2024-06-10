package complete

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.CompleteRequest) (*proto.CompleteResponse, error)

type OrderCompleter interface {
	CompleteOrderContext(ctx context.Context, id int32) error
}

func MakeCompleteHandlerFunc(log *slog.Logger, completer OrderCompleter) HandlerFunc {
	const op = "grpc.order.handler.complete.MakeCompleteHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CompleteRequest) (*proto.CompleteResponse, error) {
		log.Info("request encoded")

		err := completer.CompleteOrderContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to complete order", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to complete order: %s", err)
		}

		log.Info("order completed")

		return &proto.CompleteResponse{
			Id: req.Id,
		}, nil
	}
}
