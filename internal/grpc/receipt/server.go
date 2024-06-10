package receipt

import (
	"github.com/fatalistix/trade-organization-backend/internal/grpc/receipt/handler/createwithaccounting"
	proto "github.com/fatalistix/trade-organization-proto/gen/go/receipt"
	"google.golang.org/grpc"
	"log/slog"
)

type ServerAPI struct {
	proto.UnimplementedReceiptServiceServer
	createWithAccountingHandlerFunc createwithaccounting.HandlerFunc
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, repository *repository.Repository) {
	proto.RegisterReceiptServiceServer(
		gRPC,
		&ServerAPI{
			createWithAccountingHandlerFunc: createwithaccounting.MakeCreateWithAccountingHandlerFunc(log, repository),
		},
	)
}
