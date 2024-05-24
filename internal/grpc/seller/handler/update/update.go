package update

import (
	"context"
	modelcore "github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	protocore "github.com/fatalistix/trade-organization-backend/internal/grpc/core/mapper"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/mapper"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*emptypb.Empty, error)

type SellerUpdater interface {
	UpdateSellerContext(
		ctx context.Context,
		id int32,
		firstName *string,
		lastName *string,
		middleName *string,
		birthDate *modelcore.Date,
		salary *modelcore.Money,
		phoneNumber *string,
		worksAt *model.WorksAt,
	) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater SellerUpdater) HandlerFunc {
	const op = "grpc.seller.handler.update.MakeUpdateHandlerFunc"

	log = slog.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*emptypb.Empty, error) {
		log.Info("request encoded", slog.Any("request", req))

		var birthDate *modelcore.Date
		if req.BirthDate != nil {
			birthDateTemp := protocore.ProtoDateToModelDate(req.BirthDate)
			birthDate = &birthDateTemp
		}

		var salary *modelcore.Money
		if req.Salary != nil {
			salaryTemp := protocore.ProtoMoneyToModelMoney(req.Salary)
			salary = &salaryTemp
		}

		var worksAt *model.WorksAt
		if req.WorksAt != nil {
			var err error
			worksAt, err = mapper.ProtoWorksAtToModelWorksAt(req.WorksAt)
			if err != nil {
				log.Error("unable to convert proto place of work to model place of work", slogattr.Err(err))
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		err := updater.UpdateSellerContext(
			ctx,
			req.Id,
			req.FirstName,
			req.LastName,
			req.MiddleName,
			birthDate,
			salary,
			req.PhoneNumber,
			worksAt,
		)
		if err != nil {
			log.Error("unable to update seller", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update seller: %s", err)
		}

		return &emptypb.Empty{}, nil
	}
}
