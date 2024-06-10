package create

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
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error)

type SellerCreator interface {
	CreateSellerContext(
		ctx context.Context,
		firstName string,
		lastName string,
		middleName string,
		birthDate modelcore.Date,
		salary modelcore.Money,
		phoneNumber string,
		worksAt *model.WorksAt,
	) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator SellerCreator) HandlerFunc {
	const op = "grpc.seller.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		birthDate := protocore.ProtoDateToModelDate(req.BirthDate)
		salary := protocore.ProtoMoneyToModelMoney(req.Salary)
		worksAt, err := mapper.ProtoWorksAtToModelWorksAt(req.WorksAt)
		if err != nil {
			log.Error("unable to convert proto place of work to model place of work", slogattr.Err(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		id, err := creator.CreateSellerContext(
			ctx,
			req.FirstName,
			req.LastName,
			req.MiddleName,
			birthDate,
			salary,
			req.PhoneNumber,
			worksAt,
		)

		if err != nil {
			log.Error("unable to create seller", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to create seller: %s", err)
		}

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
