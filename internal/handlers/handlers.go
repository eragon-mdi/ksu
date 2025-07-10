package handlers

import "github.com/eragon-mdi/ksu/internal/server/routes"

type handlerImplement interface {
	routes.Handler
}

type handler struct {
	service Service
}

func New(s Service) handlerImplement {
	return &handler{
		service: s,
	}
}
