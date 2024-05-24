package seller

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/register"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/update"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/seller"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedSellerServiceServer
	registerHandlerFunc register.HandlerFunc
	listHandlerFunc     list.HandlerFunc
	updateHandlerFunc   update.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterSellerServiceServer(
		gRPC,
		&ServerAPI{
			registerHandlerFunc: register.MakeRegisterHandlerFunc(log, repository),
			listHandlerFunc:     list.MakeListHandlerFunc(log, repository),
			updateHandlerFunc:   update.MakeUpdateHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return s.registerHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) Update(ctx context.Context, req *proto.UpdateRequest) (*emptypb.Empty, error) {
	return s.updateHandlerFunc(ctx, req)
}
