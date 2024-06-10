package create

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.CreateRequest) (*proto.CreateResponse, error)

type SupplierCreator interface {
	CreateSupplierContext(
		ctx context.Context,
		name string,
	) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator SupplierCreator) HandlerFunc {
	const op = "grpc.seller.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		log.Info("creating supplier", slog.String("name", req.Name))

		id, err := creator.CreateSupplierContext(ctx, req.Name)
		if err != nil {
			log.Error("unable to create supplier", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to create supplier: %s", err)
		}

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
