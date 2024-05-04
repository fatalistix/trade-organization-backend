package seller

import (
	"context"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedSellerServiceServer
	log *slog.Logger
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger) {
	proto.RegisterSellerServiceServer(
		gRPC,
		&ServerAPI{
			log: log,
		},
	)
}

func (s *ServerAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	const op = "grpc.seller.Register"

	s.log.With(
		slog.String("op", op),
	)

	s.log.Info("request encoded", slog.Any("request", req.BirthDate))

	return &proto.RegisterResponse{}, nil
}

func (s *ServerAPI) ListByTradingPoint(ctx context.Context, req *proto.ListByTradingPointRequest) (*proto.ListByTradingPointResponse, error) {
	const op = "grpc.seller.ListByTradingPoint"

	s.log.With(
		slog.String("op", op),
	)

	s.log.Info("request encoded", slog.Any("request", req))

	return &proto.ListByTradingPointResponse{}, nil
}
