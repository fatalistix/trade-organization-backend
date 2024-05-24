package add_hall

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/hall_container"
	"github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	hallcontainermapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/hall_container"
	tradingpointmapper "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.AddHallRequest) (*proto.AddHallResponse, error)

type HallAdder interface {
	AddHallContext(
		ctx context.Context,
		hallContainerID int32,
		hallContainerType hall_container.Type,
		tradingPointID int32,
		tradingPointType trading_point.Type,
	) (int32, error)
}

func MakeAddHallHandlerFunc(log *slog.Logger, adder HallAdder) HandlerFunc {
	const op = "grpc.tradingpoint.handler.add_hall.MakeAddHallHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.AddHallRequest) (*proto.AddHallResponse, error) {
		hallContainerType, err := hallcontainermapper.ProtoTypeToModelType(req.HallContainerType)
		if err != nil {
			log.Error("unable to convert proto type to model type", slogattr.Err(err))

			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		tradingPointType, err := tradingpointmapper.ProtoTypeToModelType(req.TradingPointType)
		if err != nil {
			log.Error("unable to convert proto type to model type", slogattr.Err(err))

			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		log.Info("request encoded", slog.Any("request", req))

		id, err := adder.AddHallContext(
			ctx,
			req.HallContainerId,
			hallContainerType,
			req.TradingPointId,
			tradingPointType,
		)
		if err != nil {
			log.Error("unable to add new hall", slogattr.Err(err))

			return nil, status.Errorf(codes.Internal, err.Error())
		}

		log.Info("added new hall", slog.Any("id", id))

		return &proto.AddHallResponse{
			Id: id,
		}, nil
	}
}
