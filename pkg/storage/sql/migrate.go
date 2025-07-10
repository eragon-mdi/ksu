package sqlstorage

import (
	"github.com/eragon-mdi/ksu/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateUp(cfg config.Config, d database.Driver) error {
	cfgStor := cfg.Storage()

	m, err := migrate.NewWithDatabaseInstance(
		cfgStor.MigrateSrc(), // "file://" +
		cfgStor.DBname(),
		d,
	)
	if err != nil {
		return err
	}

	return m.Up()
}
