package list

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type CustomerProvider interface {
	CustomersContext(
		ctx context.Context,
	) ([]*proto.Customer, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider CustomerProvider) HandlerFunc {
	const op = "grpc.customer.handler.list.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		log.Info("request encoded")

		customers, err := provider.CustomersContext(ctx)
		if err != nil {
			log.Error("unable to get list of customers", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get list of customers: %s", err)
		}

		log.Info("customers", slog.Any("customers", customers))

		return &proto.ListResponse{
			Customers: customers,
		}, nil
	}
}
