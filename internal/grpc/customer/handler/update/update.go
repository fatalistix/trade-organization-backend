package update

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type CustomerUpdater interface {
	UpdateCustomerContext(ctx context.Context, customer *proto.Customer) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater CustomerUpdater) HandlerFunc {
	const op = "grpc.customer.handler.update.MakeUpdateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded")

		err := updater.UpdateCustomerContext(ctx, req.Customer)
		if err != nil {
			log.Error("unable to update customer", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update customer: %s", err)
		}

		log.Info("customer updated")

		return &proto.UpdateResponse{
			Id: req.Customer.Id,
		}, nil
	}
}
