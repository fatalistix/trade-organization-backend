package list

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type SupplierProvider interface {
	SuppliersContext(
		ctx context.Context,
	) ([]*proto.Supplier, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider SupplierProvider) HandlerFunc {
	const op = "grpc.seller.handler.list.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		suppliers, err := provider.SuppliersContext(ctx)
		if err != nil {
			log.Error("unable to get suppliers", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to get suppliers: %s", err)
		}

		log.Debug("suppliers received", slog.Any("suppliers", suppliers))

		return &proto.ListResponse{
			Suppliers: suppliers,
		}, nil
	}
}
