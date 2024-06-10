package order

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/cancel"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/complete"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/create"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/get"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/order/handler/update"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/order"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/order"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedOrderServiceServer
	createHandlerFunc   create.HandlerFunc
	listHandlerFunc     list.HandlerFunc
	getHandlerFunc      get.HandlerFunc
	updateHandlerFunc   update.HandlerFunc
	completeHandlerFunc complete.HandlerFunc
	cancelHandlerFunc   cancel.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterOrderServiceServer(
		gRPC,
		&ServerAPI{
			createHandlerFunc:   create.MakeCreateHandlerFunc(log, repository),
			listHandlerFunc:     list.MakeListHandlerFunc(log, repository),
			getHandlerFunc:      get.MakeGetHandlerFunc(log, repository),
			updateHandlerFunc:   update.MakeUpdateHandlerFunc(log, repository),
			completeHandlerFunc: complete.MakeCompleteHandlerFunc(log, repository),
			cancelHandlerFunc:   cancel.MakeCompleteHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return s.createHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) Order(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	return s.getHandlerFunc(ctx, req)
}

func (s *ServerAPI) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return s.updateHandlerFunc(ctx, req)
}

func (s *ServerAPI) Complete(ctx context.Context, req *proto.CompleteRequest) (*proto.CompleteResponse, error) {
	return nil, nil
}

func (s *ServerAPI) Cancel(ctx context.Context, req *proto.CancelRequest) (*proto.CancelResponse, error) {
	return nil, nil
}
