package create

import (
	"context"
	protocore "github.com/fatalistix/trade-organization-proto/gen/go/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.CreateRequest) (*proto.CreateResponse, error)

type CustomerCreator interface {
	CreateCustomerContext(
		ctx context.Context,
		firstName string,
		lastName string,
		birthDate *protocore.Date,
		phoneNumber string,
	) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator CustomerCreator) HandlerFunc {
	const op = "grpc.customer.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		log.Info("create customer request", slog.String("first_name", req.FirstName), slog.String("last_name", req.LastName), slog.String("phone_number", req.PhoneNumber))

		id, err := creator.CreateCustomerContext(
			ctx, req.FirstName, req.LastName, req.BirthDate, req.PhoneNumber,
		)
		if err != nil {
			log.Error("unable to create customer", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to create customer: %s", err)
		}

		log.Info("customer created", slog.Int("id", int(id)))

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
