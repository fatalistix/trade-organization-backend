package customer

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/customer/handler/create"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/customer/handler/delete"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/customer/handler/get"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/customer/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/customer/handler/update"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/customer"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/customer"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedCustomerServiceServer
	createHandlerFunc create.HandlerFunc
	listHandlerFunc   list.HandlerFunc
	getHandlerFunc    get.HandlerFunc
	updateHandlerFunc update.HandlerFunc
	deleteHandlerFunc delete.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterCustomerServiceServer(
		gRPC,
		&ServerAPI{
			createHandlerFunc: create.MakeCreateHandlerFunc(log, repository),
			listHandlerFunc:   list.MakeListHandlerFunc(log, repository),
			getHandlerFunc:    get.MakeGetHandlerFunc(log, repository),
			updateHandlerFunc: update.MakeUpdateHandlerFunc(log, repository),
			deleteHandlerFunc: delete.MakeDeleteHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return s.createHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) Customer(ctx context.Context, req *proto.CustomerRequest) (*proto.CustomerResponse, error) {
	return s.getHandlerFunc(ctx, req)
}

func (s *ServerAPI) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return s.updateHandlerFunc(ctx, req)
}

func (s *ServerAPI) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return s.deleteHandlerFunc(ctx, req)
}
