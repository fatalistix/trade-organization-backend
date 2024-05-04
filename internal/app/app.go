package app

import (
	"fmt"
	grpcapp "github.com/fatalistix/trade-organization-backend/internal/app/grpc"
	"github.com/fatalistix/trade-organization-backend/internal/config"
	pgconnection "github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	pgmigration "github.com/fatalistix/trade-organization-backend/internal/database/migration/postgres"
	"log/slog"
	"os"
)

type App struct {
	grpcApp  *grpcapp.App
	migrator *pgmigration.Migrator
	database *pgconnection.Database
}

func NewApp(log *slog.Logger, cfg *config.Config) (*App, error) {
	const op = "app.NewApp"

	database, err := pgconnection.NewDatabase(
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.User,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.DBName,
		cfg.PostgreSQL.SSLMode,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to connect to database: %w", op, err)
	}

	log.Info("connected to database")

	migrationFiles := os.DirFS(cfg.Migrations.Path)

	migrator, err := pgmigration.NewMigrator(migrationFiles, ".")
	if err != nil {
		return nil, fmt.Errorf("%s: unable to create migrator: %w", op, err)
	}

	log.Info("migrator created")

	err = migrator.ApplyMigrations(database.DB().DB, cfg.PostgreSQL.DBName)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to apply migrations: %w", op, err)
	}

	log.Info("migrations applied")

	grpcApp := grpcapp.NewApp(log, int(cfg.GRPC.Port), database)

	return &App{grpcApp: grpcApp, database: database, migrator: migrator}, nil
}

func (a *App) Run() error {
	return a.grpcApp.Run()
}

func (a *App) Stop() {
	defer func() {
		_ = a.database.DB().Close()
		_ = a.migrator.Close()
	}()

	a.grpcApp.Stop()
}
