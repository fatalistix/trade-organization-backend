package list

import (
	"context"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/trading_point"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/mapper/trading_point"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type TradingPointProvider interface {
	TradingPointsContext(ctx context.Context) ([]model.TradingPoint, error)
}

type HandlerFunc = func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

func MakeListHandlerFunc(log *slog.Logger, provider TradingPointProvider) HandlerFunc {
	const op = "grpc.register.List"

	log = log.With(
		slog.String("op", op),
	)

	return func(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error) {
		tps, err := provider.TradingPointsContext(ctx)
		if err != nil {
			log.Error("unable to get list of all points", slogattr.Err(err))

			return nil, status.Errorf(codes.Internal, err.Error())
		}

		log.Info("received list of all trading points", slog.Any("trading_points", tps))

		protoTPs := make([]*proto.TradingPoint, 0, len(tps))
		for _, tp := range tps {
			ptp, err := trading_point.ModelTradingPointToProtoTradingPoint(tp)
			if err != nil {
				log.Error("error mapping from model type to proto type")

				return nil, status.Errorf(codes.Internal, err.Error())
			}
			protoTPs = append(protoTPs, ptp)
		}

		return &proto.ListResponse{
			TradingPoints: protoTPs,
		}, nil
	}
}
