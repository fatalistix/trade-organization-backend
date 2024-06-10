package seller

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/create"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/get"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/handler/patch"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/seller"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedSellerServiceServer
	createHandlerFunc create.HandlerFunc
	listHandlerFunc   list.HandlerFunc
	getHandlerFunc    get.HandlerFunc
	patchHandlerFunc  patch.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterSellerServiceServer(
		gRPC,
		&ServerAPI{
			createHandlerFunc: create.MakeCreateHandlerFunc(log, repository),
			listHandlerFunc:   list.MakeListHandlerFunc(log, repository),
			getHandlerFunc:    get.MakeGetHandlerFunc(log, repository),
			patchHandlerFunc:  patch.MakePatchHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return s.createHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) Patch(ctx context.Context, req *proto.PatchRequest) (*proto.PatchResponse, error) {
	return s.patchHandlerFunc(ctx, req)
}

func (s *ServerAPI) Seller(ctx context.Context, req *proto.SellerRequest) (*proto.SellerResponse, error) {
	return s.getHandlerFunc(ctx, req)
}
