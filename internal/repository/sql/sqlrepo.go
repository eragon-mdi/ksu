package sqlrepo

import "database/sql"

type SQLStorage interface {
	SQLDB() *sql.DB
}

type sqlRepository struct {
	storage SQLStorage
}

func New(s SQLStorage) *sqlRepository {
	return &sqlRepository{
		storage: s,
	}
}
