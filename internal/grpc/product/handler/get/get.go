package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.ProductRequest) (*proto.ProductResponse, error)

type ProductProvider interface {
	ProductContext(ctx context.Context, id int32) (*proto.Product, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider ProductProvider) HandlerFunc {
	const op = "grpc.product.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		product, err := provider.ProductContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get product", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get product: %s", err)
		}

		log.Info("product received", slog.Any("product", product))

		return &proto.ProductResponse{
			Product: product,
		}, nil
	}
}
