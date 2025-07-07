package sqlrepo

import (
	"database/sql"
)

type SQLStorage struct {
	*sql.DB
}

type sqlRepository struct {
	storage SQLStorage
}

func New(s SQLStorage) *sqlRepository {
	return &sqlRepository{
		storage: s,
	}
}
