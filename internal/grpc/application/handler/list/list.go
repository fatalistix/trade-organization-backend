package list

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc func(context.Context, *proto.ListRequest) (*proto.ListResponse, error)

type ApplicationProvider interface {
	ApplicationsContext(
		ctx context.Context,
	) ([]*proto.Application, error)
}

func MakeListHandlerFunc(log *slog.Logger, provider ApplicationProvider) HandlerFunc {
	const op = "grpc.application.handler.list.MakeListHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
		applications, err := provider.ApplicationsContext(ctx)
		if err != nil {
			log.Error("unable to get applications", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get applications: %s", err)
		}

		log.Debug("applications received", slog.Any("applications", applications))

		return &proto.ListResponse{
			Applications: applications,
		}, nil
	}
}
