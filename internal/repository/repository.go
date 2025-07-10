package repository

import (
	"fmt"

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
	fmt.Printf("got type: %T\n", v)

	switch s := v.(type) {
	case sqlrepo.SQLStorage:
		return sqlrepo.New(s)
	case fakerepo.FakeStorage:
		return fakerepo.New(s)
	default:
		panic(fmt.Sprintf("repository: undefined storage type: %T\n", v))
	}
}
