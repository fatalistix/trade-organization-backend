package get

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(context.Context, *proto.ApplicationRequest) (*proto.ApplicationResponse, error)

type ApplicationProvider interface {
	ApplicationContext(ctx context.Context, id int32) (*proto.Application, error)
}

func MakeGetHandlerFunc(log *slog.Logger, provider ApplicationProvider) HandlerFunc {
	const op = "grpc.application.handler.get.MakeGetHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.ApplicationRequest) (*proto.ApplicationResponse, error) {
		log.Info("request encoded")

		application, err := provider.ApplicationContext(ctx, req.Id)
		if err != nil {
			log.Error("unable to get application", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to get application: %s", err)
		}

		log.Info("application received", slog.Any("application", application))

		return &proto.ApplicationResponse{
			Application: application,
		}, nil
	}
}
