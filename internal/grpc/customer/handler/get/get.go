package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.CustomerRequest) (*proto.CustomerResponse, error)

type CustomerProvider interface {
	CustomerContext(ctx context.Context, id int32) (*proto.Customer, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider CustomerProvider) HandlerFunc {
	const op = "grpc.customer.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CustomerRequest) (*proto.CustomerResponse, error) {
		log.Info("request encoded")

		customer, err := provider.CustomerContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get customer", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get customer: %s", err)
		}

		log.Info("customer received")

		return &proto.CustomerResponse{
			Customer: customer,
		}, nil
	}
}
