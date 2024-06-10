package patch

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	protocore "github.com/fatalistix/trade-organization-proto/gen/go/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type SellerUpdater interface {
	UpdateSellerContext(
		ctx context.Context,
		id int32,
		firstName *string,
		lastName *string,
		middleName *string,
		birthDate *protocore.Date,
		salary *protocore.Money,
		phoneNumber *string,
		worksAt *proto.NewWorksAt,
	) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater SellerUpdater) HandlerFunc {
	const op = "grpc.seller.handler.update.MakeUpdateHandlerFunc"

	log = slog.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		err := updater.UpdateSellerContext(
			ctx,
			req.Id,
			req.FirstName,
			req.LastName,
			req.MiddleName,
			req.BirthDate,
			req.Salary,
			req.PhoneNumber,
			req.WorksAt,
		)
		if err != nil {
			log.Error("unable to update seller", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update seller: %s", err)
		}

		return &proto.UpdateResponse{
			Id: req.Id,
		}, nil
	}
}
