package register

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/core"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	grpccore "github.com/fatalistix/trade-organization-backend/internal/grpc/core/mapper"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)

type TradingPointRegisterer interface {
	RegisterTradingPoint(
		ctx context.Context,
		t model.Type,
		areaPlot float64,
		rentalCharge core.Money,
		counterCount int32,
		address string,
	) (int32, error)
}

func MakeRegisterHandlerFunc(log *slog.Logger, registerer TradingPointRegisterer) HandlerFunc {
	const op = "grpc.tradingpoint.handler.register.MakeRegisterHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
		t, err := trading_point.ProtoTypeToModelType(req.Type)
		if err != nil {
			log.Error("unable to convert proto type to model type", slogattr.Err(err))

			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if req.RentalCharge.Pennies > 99 {
			log.Error("invalid rental charge", slog.Any("rental charge", req.RentalCharge))

			return nil, status.Errorf(codes.InvalidArgument, "invalid rental charge")
		}

		log.Info("request encoded", slog.Any("request", req))

		id, err := registerer.RegisterTradingPoint(
			ctx,
			t,
			req.AreaPlot,
			grpccore.ProtoMoneyToModelMoney(req.RentalCharge),
			req.CounterCount,
			req.Address,
		)
		if err != nil {
			log.Error("unable to register new trading point", slogattr.Err(err))

			return nil, status.Errorf(codes.Internal, err.Error())
		}

		log.Info("registered new trading point", slog.Any("id", id), slog.Any("type", t))

		return &proto.RegisterResponse{Id: id}, nil
	}
}
