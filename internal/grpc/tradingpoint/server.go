package tradingpoint

import (
	"context"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/add_hall"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/add_section"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/department_store"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/gettradingpoint"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/list"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint/handler/register"
	repository "github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedTradingPointServiceServer
	addHallHandlerFunc         add_hall.HandlerFunc
	addSectionHandlerFunc      add_section.HandlerFunc
	departmentStoreHandlerFunc department_store.HandlerFunc
	registerHandlerFunc        register.HandlerFunc
	listHandlerFunc            list.HandlerFunc
	getTradingPointFunc        gettradingpoint.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterTradingPointServiceServer(
		gRPC,
		&ServerAPI{
			addHallHandlerFunc:         add_hall.MakeAddHallHandlerFunc(log, repository),
			addSectionHandlerFunc:      add_section.MakeAddSectionHandlerFunc(log, repository),
			departmentStoreHandlerFunc: department_store.MakeDepartmentStoreHandlerFunc(log, repository),
			registerHandlerFunc:        register.MakeRegisterHandlerFunc(log, repository),
			listHandlerFunc:            list.MakeListHandlerFunc(log, repository),
			getTradingPointFunc:        gettradingpoint.MakeGetHandlerFunc(log, repository),
		},
	)
}

func (s *ServerAPI) AddHall(ctx context.Context, req *proto.AddHallRequest) (*proto.AddHallResponse, error) {
	return s.addHallHandlerFunc(ctx, req)
}

func (s *ServerAPI) AddSection(ctx context.Context, req *proto.AddSectionRequest) (*proto.AddSectionResponse, error) {
	return s.addSectionHandlerFunc(ctx, req)
}

func (s *ServerAPI) DepartmentStore(ctx context.Context, req *proto.DepartmentStoreRequest) (*proto.DepartmentStoreResponse, error) {
	return s.departmentStoreHandlerFunc(ctx, req)
}

func (s *ServerAPI) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return s.registerHandlerFunc(ctx, req)
}

func (s *ServerAPI) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	return s.listHandlerFunc(ctx, req)
}

func (s *ServerAPI) TradingPoint(ctx context.Context, req *proto.TradingPointRequest) (*proto.TradingPointResponse, error) {
	return s.getTradingPointFunc(ctx, req)
}
