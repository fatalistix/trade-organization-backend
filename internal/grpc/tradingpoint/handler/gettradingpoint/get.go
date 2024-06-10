package gettradingpoint

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.TradingPointRequest) (*proto.TradingPointResponse, error)

type TradingPointProvider interface {
	TradingPointContext(ctx context.Context, id int32, t proto.TradingPointType) (*proto.TradingPoint, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider TradingPointProvider) HandlerFunc {
	const op = "grpc.tradingpoint.handler.gettradingpoint.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.TradingPointRequest) (*proto.TradingPointResponse, error) {
		log.Debug("Get trading point", slog.Int("id", int(req.Id)))

		tradingPoint, err := provider.TradingPointContext(ctx, req.Id, req.Type)
		if err != nil {
			return nil, err
		}

		return &proto.TradingPointResponse{
			TradingPoint: tradingPoint,
		}, nil
	}
}
