package update

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type ProductUpdater interface {
	UpdateProductContext(ctx context.Context, product *proto.Product) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater ProductUpdater) HandlerFunc {
	const op = "grpc.product.handler.update.MakeUpdateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		err := updater.UpdateProductContext(ctx, req.Product)
		if err != nil {
			log.Error("unable to update product", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update product: %s", err)
		}

		log.Info("product updated", slog.Any("product", req.Product))

		return &proto.UpdateResponse{
			Id: req.Product.Id,
		}, nil
	}
}
