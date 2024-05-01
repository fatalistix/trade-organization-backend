package main

import (
	"github.com/fatalistix/trade-organization-backend/internal/app"
	"github.com/fatalistix/trade-organization-backend/internal/config"
	"github.com/fatalistix/trade-organization-backend/internal/env"
	slogattr "github.com/fatalistix/trade-organization-backend/internal/lib/log/slog/attr"
	"github.com/golang-cz/devslog"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := setupLogger()

	env.MustLoadEnv()
	log.Info("environment variables loaded")

	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.MustLoadConfig(configPath)
	log.Info("config loaded")
	log.Info("server info:", slog.Int("port", int(cfg.GRPC.Port)))
	log.Debug("config", slog.Any("config", cfg))

	a, err := app.NewApp(log, cfg)
	if err != nil {
		log.Error("unable to create app", slogattr.Err(err))
		os.Exit(1)
	}

	log.Info("starting application")

	go func() {
		if err := a.Run(); err != nil {
			log.Error("unable to run app", slogattr.Err(err))
		}
	}()

	log.Info("application started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	a.Stop()
	log.Info("application stopped")
}

func setupLogger() *slog.Logger {
	slogOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	opts := &devslog.Options{
		HandlerOptions:  slogOpts,
		SortKeys:        true,
		NewLineAfterLog: true,
	}

	return slog.New(devslog.NewHandler(os.Stdout, opts))
}
