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
	//var postrgres *postrgres

	//register("postgres", &InitFuncs{
	//	Connect: connectPostgresAdapter,
	//	Migrate: sqlstorage.MigratePostgres,
	//})
	register("postgres", &postrgres{})
}

//	func connectPostgresAdapter(cfg config.Config) (storageImplement, error) {
//		return sqlstorage.ConnectPostgres(cfg)
//	}

func (s *postrgres) Connect(cfg config.Config) (err error) {
	s.DB, err = sqlstorage.ConnectPostgres(cfg)
	return
}

func (s *postrgres) Migrate(cfg config.Config) error {
	return sqlstorage.MigratePostgres(cfg, s.DB)
}
