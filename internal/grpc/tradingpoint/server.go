package tradingpoint

import (
	"context"
	"fmt"
	grpccore "github.com/fatalistix/trade-organization-backend/internal/grpc/core"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	"github.com/fatalistix/trade-organization-backend/internal/model/core"
	model "github.com/fatalistix/trade-organization-backend/internal/model/tradingpoint"
	protocore "github.com/fatalistix/trade-organization-proto/gen/go/core"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Provider interface {
	RegisterNewTradingPoint(
		ctx context.Context,
		t model.Type,
		areaPlot float64,
		rentalCharge *core.Money,
		counterCount int32,
		address string,
	) (int32, error)
	List(ctx context.Context) ([]*model.TradingPoint, error)
}

type ServerAPI struct {
	proto.UnimplementedTradingPointServiceServer
	provider Provider
	log      *slog.Logger
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, provider Provider) {
	proto.RegisterTradingPointServiceServer(
		gRPC,
		&ServerAPI{
			provider: provider,
			log:      log,
		},
	)
}

func (s *ServerAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	const op = "grpc.tradingpoint.Register"

	s.log.With(
		slog.String("op", op),
	)

	t, err := ProtoTypeToModelType(req.Type)
	if err != nil {
		s.log.Error("unable to convert proto type to model type", slogattr.Err(err))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if req.RentalCharge.Pennies > 99 {
		s.log.Error("invalid rental charge", slog.Any("rental charge", req.RentalCharge))

		return nil, status.Errorf(codes.InvalidArgument, "invalid rental charge")
	}

	s.log.Info("request encoded", slog.Any("request", req))

	id, err := s.provider.RegisterNewTradingPoint(
		ctx,
		t,
		req.AreaPlot,
		grpccore.ProtoMoneyToProtoMoney(req.RentalCharge),
		req.CounterCount,
		req.Address,
	)
	if err != nil {
		s.log.Error("unable to register new trading point", slogattr.Err(err))

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.log.Info("registered new trading point", slog.Any("id", id), slog.Any("type", t))

	return &proto.RegisterResponse{Id: id}, nil
}

func (s *ServerAPI) List(ctx context.Context, _ *protocore.Empty) (*proto.ListResponse, error) {
	const op = "grpc.register.List"

	// TODO: replace with `log = s.log.With...`
	s.log.With(
		slog.String("op", op),
	)

	tps, err := s.provider.List(ctx)
	if err != nil {
		s.log.Error("unable to get list of all points", slogattr.Err(err))

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.log.Info("received list of all trading points", slog.Any("trading_points", tps))

	protoTPs := make([]*proto.TradingPoint, 0, len(tps))
	for _, tp := range tps {
		ptp, err := modelTradingPointToProtoTradingPoint(tp)
		if err != nil {
			s.log.Error("error mapping from model type to proto type")

			return nil, status.Errorf(codes.Internal, err.Error())
		}
		protoTPs = append(protoTPs, ptp)
	}

	return &proto.ListResponse{
		TradingPoints: protoTPs,
	}, nil
}

func modelTradingPointToProtoTradingPoint(tp *model.TradingPoint) (*proto.TradingPoint, error) {
	const op = "grpc.register.modelTradingPointToProtoTradingPoint"

	protoType, err := ModelTypeToProtoType(tp.Type)
	if err != nil {
		return nil, fmt.Errorf("%s: error mapping: %w", op, err)
	}
	return &proto.TradingPoint{
		Id:           tp.ID,
		Type:         protoType,
		AreaPlot:     tp.AreaPlot,
		RentalCharge: grpccore.ModelMoneyToProtoMoney(tp.RentalCharge),
		CounterCount: tp.CounterCount,
		Address:      tp.Address,
	}, nil
}
