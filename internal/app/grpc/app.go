package grpc

import (
	"context"
	"fmt"
	"github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	applicationgrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/application"
	customergrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/customer"
	ordergrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/order"
	productgrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/product"
	sellergrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/seller"
	suppliergrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/supplier"
	tradingpointgrpc "github.com/fatalistix/trade-organization-backend/internal/grpc/tradingpoint"
	applicationrepository "github.com/fatalistix/trade-organization-backend/internal/repository/application"
	customerrepository "github.com/fatalistix/trade-organization-backend/internal/repository/customer"
	orderrepository "github.com/fatalistix/trade-organization-backend/internal/repository/order"
	productrepository "github.com/fatalistix/trade-organization-backend/internal/repository/product"
	sellerrepository "github.com/fatalistix/trade-organization-backend/internal/repository/seller"
	supplierrepository "github.com/fatalistix/trade-organization-backend/internal/repository/supplier"
	tradingpointrepository "github.com/fatalistix/trade-organization-backend/internal/repository/tradingpoint"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
)

type App struct {
	port       int
	grpcServer *grpc.Server
}

func NewApp(log *slog.Logger, port int, database *postgres.Database) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.StartCall,
			logging.FinishCall,
		),
	}
	_ = loggingOpts

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	_ = recoveryOpts

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
		),
	)

	productRepository := productrepository.NewRepository(database)

	productgrpc.RegisterServer(grpcServer, log, productRepository)
	tradingpointgrpc.RegisterServer(grpcServer, log, tradingpointrepository.NewRepository(database, productRepository))
	sellergrpc.RegisterServer(grpcServer, log, sellerrepository.NewRepository(database))
	suppliergrpc.RegisterServer(grpcServer, log, supplierrepository.NewRepository(database, productRepository))
	applicationgrpc.RegisterServer(grpcServer, log, applicationrepository.NewRepository(database))
	customergrpc.RegisterServer(grpcServer, log, customerrepository.NewRepository(database))
	ordergrpc.RegisterServer(grpcServer, log, orderrepository.NewRepository(database))

	return &App{
		port:       port,
		grpcServer: grpcServer,
	}
}

// InterceptorLogger adapts slog logger to interceptor logger
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

// Run starts gRPC server
func (a *App) Run() error {
	const op = "app.grpc.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server
func (a *App) Stop() {
	a.grpcServer.GracefulStop()
}
