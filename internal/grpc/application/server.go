package application

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/application/handler/create"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/application/handler/get"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/application/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/application/handler/update"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/application"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedApplicationServiceServer
	createHandlerFunc create.HandlerFunc
	listHandlerFunc   list.HandlerFunc
	getHandlerFunc    get.HandlerFunc
	updateHandlerFunc update.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterApplicationServiceServer(
		gRPC,
		&ServerAPI{
			createHandlerFunc: create.MakeCreateHandlerFunc(log, repository),
			listHandlerFunc:   list.MakeListHandlerFunc(log, repository),
			getHandlerFunc:    get.MakeGetHandlerFunc(log, repository),
			updateHandlerFunc: update.MakeUpdateHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return s.createHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) Application(ctx context.Context, req *proto.ApplicationRequest) (*proto.ApplicationResponse, error) {
	return s.getHandlerFunc(ctx, req)
}

func (s *ServerAPI) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return s.updateHandlerFunc(ctx, req)
}
