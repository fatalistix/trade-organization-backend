package add_section

import (
	"context"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/tradingpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type HandlerFunc = func(ctx context.Context, req *proto.AddSectionRequest) (*proto.AddSectionResponse, error)

type SectionAdder interface {
	AddSectionContext(
		ctx context.Context,
		departmentStoreID int32,
	) (int32, error)
}

func MakeAddSectionHandlerFunc(log *slog.Logger, adder SectionAdder) HandlerFunc {
	const op = "grpc.tradingpoint.handler.add_section.MakeAddSectionHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(ctx context.Context, req *proto.AddSectionRequest) (*proto.AddSectionResponse, error) {
		id, err := adder.AddSectionContext(ctx, req.DepartmentStoreId)
		if err != nil {
			log.Error("unable to add new section", slogattr.Err(err))

			return nil, status.Errorf(codes.Internal, err.Error())
		}

		log.Info("added new section", slog.Any("id", id))

		return &proto.AddSectionResponse{
			Id: id,
		}, nil
	}
}
