package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type sslMode string

const (
	SSLDisable    sslMode = "disable"
	SSLRequire    sslMode = "require"
	SSLVerifyFull sslMode = "verify-full"
	SSLVerifyCa   sslMode = "verify-ca"
)

func NewDB(host string, port int, user string, password string, dbname string, sslMode sslMode) (*sql.DB, error) {
	const op = "database.connection.postgres.NewDB"

	pgCredentials := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode,
	)

	db, err := sql.Open("postgres", pgCredentials)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// verify that data source name is valid (according to godoc of `sql.Open`)
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
