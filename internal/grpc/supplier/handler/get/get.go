package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.SupplierRequest) (*proto.SupplierResponse, error)

type SupplierProvider interface {
	SupplierContext(ctx context.Context, id int32) (*proto.Supplier, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider SupplierProvider) HandlerFunc {
	const op = "grpc.supplier.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.SupplierRequest) (*proto.SupplierResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		supplier, err := provider.SupplierContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get supplier", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get supplier: %s", err)
		}

		log.Info("supplier received", slog.Any("supplier", supplier))

		return &proto.SupplierResponse{
			Supplier: supplier,
		}, nil
	}
}
