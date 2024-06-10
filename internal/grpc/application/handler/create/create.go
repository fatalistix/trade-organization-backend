package create

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	prototradingpoint "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error)

type ApplicationCreator interface {
	CreateApplicationContext(
		ctx context.Context,
		tradingPointID int32,
		tradingPointType prototradingpoint.TradingPointType,
		products []*proto.ProductApplication,
	) (int32, error)
}

func MakeCreateHandlerFunc(log *slog.Logger, creator ApplicationCreator) HandlerFunc {
	const op = "grpc.application.handler.create.MakeCreateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
		log.Info("request encoded")

		id, err := creator.CreateApplicationContext(ctx, req.TradingPointId, prototradingpoint.TradingPointType(req.TradingPointType), req.Products)
		if err != nil {
			log.Error("unable to create application", slog.Any("err", err))
			return nil, status.Errorf(codes.Internal, "unable to create application: %s", err)
		}

		log.Info("application created")

		return &proto.CreateResponse{
			Id: id,
		}, nil
	}
}
