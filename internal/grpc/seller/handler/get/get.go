package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/seller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.SellerRequest) (*proto.SellerResponse, error)

type SellerProvider interface {
	SellerContext(ctx context.Context, id int32) (*proto.Seller, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider SellerProvider) HandlerFunc {
	const op = "grpc.seller.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.SellerRequest) (*proto.SellerResponse, error) {
		log.Info("request encoded")

		seller, err := provider.SellerContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get seller", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get seller: %s", err)
		}

		log.Info("seller received")

		return &proto.SellerResponse{
			Seller: seller,
		}, nil
	}
}
