package service

import "github.com/eragon-mdi/ksu/internal/repository"

type Servicer interface {
}

type service struct {
	repository repository.Repositorier
}

func New(r repository.Repositorier) Servicer {
	return service{
		repository: r,
	}
}
