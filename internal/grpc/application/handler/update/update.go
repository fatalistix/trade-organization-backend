package update

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error)

type ApplicationUpdater interface {
	UpdateApplicationContext(ctx context.Context, application *proto.Application) error
}

func MakeUpdateHandlerFunc(log *slog.Logger, updater ApplicationUpdater) HandlerFunc {
	const op = "grpc.application.handler.update.MakeUpdateHandlerFunc"

	log = log.With(
		log, slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
		log.Info("request encoded")

		err := updater.UpdateApplicationContext(ctx, req.Application)
		if err != nil {
			log.Error("unable to update application", slogattr.Err(err))
			return nil, status.Errorf(codes.Internal, "unable to update application: %s", err)
		}

		log.Info("application updated")

		return &proto.UpdateResponse{
			Id: req.Application.Id,
		}, nil
	}
}
