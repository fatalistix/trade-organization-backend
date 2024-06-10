package update

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/supplier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type SupplierUpdater interface {
	UpdateSupplierContext(ctx context.Context, supplier *proto.Supplier) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater SupplierUpdater) HandlerFunc {
	const op = "grpc.supplier.handler.update.MakeUpdateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded")

		err := updater.UpdateSupplierContext(ctx, req.Supplier)
		if err != nil {
			log.Error("unable to update supplier", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update supplier: %s", err)
		}

		log.Info("supplier updated")

		return &proto.UpdateResponse{
			Id: req.Supplier.Id,
		}, nil
	}
}
