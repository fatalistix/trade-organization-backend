package create

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.CreateRequest) (*proto.CreateResponse, error)

type ProductCreator interface {
	CreateProductContext(ctx context.Context, name, description string) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator ProductCreator) HandlerFunc {
	const op = "grpc.product.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		id, err := creator.CreateProductContext(ctx, req.Name, req.Description)
		if err != nil {
			log.Error("unable to create product", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to create product: %s", err)
		}

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
