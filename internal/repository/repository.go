package repository

import (
	fakerepo "github.com/eragon-mdi/ksu/internal/repository/fake"
	sqlrepo "github.com/eragon-mdi/ksu/internal/repository/sql"
	"github.com/eragon-mdi/ksu/internal/service"
	taskstate "github.com/eragon-mdi/ksu/internal/service/task_state"
)

type Repository interface {
	service.Repository
	taskstate.Repository
}

// скастить тип SQL | FAKE | NOSQL ...
type Storage interface{}

func New(v Storage) Repository {
	switch s := v.(type) {
	case sqlrepo.SQLStorage:
		return sqlrepo.New(s)
	case fakerepo.FakeStorage:
		return fakerepo.New(s)
	default:
		panic("repository: undefined storage")
	}
}
