package delete

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.DeleteRequest) (*proto.DeleteResponse, error)

type CustomerDeleter interface {
	DeleteCustomerContext(ctx context.Context, id int32) error
}

func MakeDeleteHandlerFunc(log *slog.Logger, deleter CustomerDeleter) HandlerFunc {
	const op = "grpc.customer.handler.delete.MakeDeleteHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
		log.Info("delete customer request", slog.Int("id", int(req.Id)))

		err := deleter.DeleteCustomerContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to delete customer", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to delete customer: %s", err)
		}

		log.Info("customer deleted", slog.Int("id", int(req.Id)))

		return &proto.DeleteResponse{
			Id: req.Id,
		}, nil
	}
}
