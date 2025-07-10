package storage

import (
	"database/sql"

	"github.com/eragon-mdi/ksu/pkg/config"
	sqlstorage "github.com/eragon-mdi/ksu/pkg/storage/sql"
	_ "github.com/lib/pq"
)

type postrgres struct {
	*sql.DB
}

func init() {
	register("postgres", &postrgres{})
}

func (s *postrgres) Connect(cfg config.Config) (err error) {
	s.DB, err = sqlstorage.ConnectPostgres(cfg)
	return
}

func (s *postrgres) Migrate(cfg config.Config) error {
	return sqlstorage.MigratePostgres(cfg, s.DB)
}

// need for switch type in repo
func (s *postrgres) SQLDB() *sql.DB {
	return s.DB
}

func (s *postrgres) Shutdown() error {
	return s.Close()
}
