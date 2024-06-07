package list

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type ProductProvider interface {
	ProductsContext(ctx context.Context, ids *[]int32) ([]*proto.Product, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider ProductProvider) HandlerFunc {
	const op = "grpc.product.handler.list.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		products, err := provider.ProductsContext(ctx, nil)
		if err != nil {
			log.Error("unable to list products", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to list products: %s", err)
		}

		return &proto.ListResponse{
			Products: products,
		}, nil
	}
}
