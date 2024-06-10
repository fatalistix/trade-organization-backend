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

type HandlerFunc = func(ctx context.Context, req *proto.PatchRequest) (*proto.PatchResponse, error)

type SellerPatcher interface {
	PatchSellerContext(
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

func MakePatchHandlerFunc(log *slog.Logger, patcher SellerPatcher) HandlerFunc {
	const op = "grpc.seller.handler.patch.MakePatchHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.PatchRequest) (*proto.PatchResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		err := patcher.PatchSellerContext(
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
			log.Error("unable to patch seller", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to patch seller: %s", err)
		}

		return &proto.PatchResponse{
			Id: req.Id,
		}, nil
	}
}
