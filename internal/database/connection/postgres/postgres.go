package postgres

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/lib/pq"
)

type Database struct {
	db *bun.DB
}

func NewDatabase(host string, port uint16, user string, password string, dbname string, sslMode string) (*Database, error) {
	const op = "database.connection.postgres.NewDB"

	pgCredentials := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode,
	)

	sqlDB, err := sql.Open("postgres", pgCredentials)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// verify that data source name is valid (according to godoc of `sql.Open`)
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	return &Database{db: bunDB}, nil
}

func (d *Database) DB() *bun.DB {
	return d.db
}
