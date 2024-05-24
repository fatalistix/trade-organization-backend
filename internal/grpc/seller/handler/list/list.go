package list

import (
	"context"
	model "github.com/fatalistix/trade-organization-backend/internal/domain/model/seller"
	"github.com/fatalistix/trade-organization-backend/internal/grpc/seller/mapper"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type SellerProvider interface {
	SellersContext(
		ctx context.Context, filter *model.Filter,
	) ([]model.Seller, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider SellerProvider) HandlerFunc {
	const op = "grpc.seller.handler.list_by_trading_point.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		filter, err := mapper.ListRequestToModelFilter(req)
		if err != nil {
			log.Error("unable to convert list request to model filter", slogattr.Err(err))
			return nil, status.Errorf(codes.InvalidArgument, "unable to convert list request to model filter: %s", err)
		}

		log.Info("request encoded", slog.Any("filter", filter))

		sellers, err := provider.SellersContext(
			ctx, filter,
		)
		if err != nil {
			log.Error("unable to get list of sellers", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get list of sellers: %s", err)
		}

		log.Info("sellers", slog.Any("sellers", sellers))

		protoSellers := make([]*proto.Seller, 0, len(sellers))
		for _, s := range sellers {
			ps, err := mapper.ModelSellerToProtoSeller(s)
			if err != nil {
				log.Error("error mapping from model type to proto type", slogattr.Err(err))
				return nil, status.Errorf(codes.Internal, "unable to get list : %s", err)
			}

			protoSellers = append(protoSellers, ps)
		}

		return &proto.ListResponse{
			Sellers: protoSellers,
		}, nil
	}
}
