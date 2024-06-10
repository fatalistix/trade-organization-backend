package department_store

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.DepartmentStoreRequest) (*proto.DepartmentStoreResponse, error)

type DepartmentStoreProvider interface {
	DepartmentStoreContext(ctx context.Context, id int32) (*proto.DepartmentStore, error)
}

func MakeDepartmentStoreHandlerFunc(log *slog.Logger, provider DepartmentStoreProvider) HandlerFunc {
	const op = "grpc.tradingpoint.handler.department_store.MakeDepartmentStoreHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.DepartmentStoreRequest) (*proto.DepartmentStoreResponse, error) {
		log.Info("request encoded", slog.Any("request", req))

		ds, err := provider.DepartmentStoreContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get department store", slogattr.Err(err))

			return nil, status.Errorf(codes.Internal, err.Error())
		}

		log.Info("received department store", slog.Any("department_store", ds))

		//sections, err := provider.SectionsContext(ctx, &req.Id)
		//if err != nil {
		//	log.Error("unable to get sections", slogattr.Err(err))
		//
		//	return nil, status.Errorf(codes.Internal, err.Error())
		//}
		//
		//log.Info("received sections", slog.Any("sections", sections))
		//
		//tradingPointType := trading_point.TypeDepartmentStore
		//halls, err := provider.HallsContext(ctx, &req.Id, &tradingPointType)
		//if err != nil {
		//	log.Error("unable to get halls", slogattr.Err(err))
		//
		//	return nil, status.Errorf(codes.Internal, err.Error())
		//}
		//
		//log.Info("received halls", slog.Any("halls", halls))

		return &proto.DepartmentStoreResponse{
			DepartmentStore: ds,
		}, nil
	}
}
