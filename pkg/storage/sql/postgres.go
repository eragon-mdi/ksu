package sqlstorage

import (
	"database/sql"
	"fmt"

	"github.com/eragon-mdi/ksu/pkg/config"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.Config) (*sql.DB, error) {
	cfgDb := cfg.Storage()

	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfgDb.Host(), cfgDb.Port(), cfgDb.User(), cfgDb.Password(), cfgDb.DBname(), cfgDb.SSLmode(),
	)

	conn, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func MigratePostgres(cfg config.Config, conn *sql.DB) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return err
	}

	if err := migrateUp(cfg, driver); err != nil {
		return err
	}

	return nil
}
