package handlers

import "github.com/eragon-mdi/ksu/internal/service"

type Handler interface {
}

type handler struct {
	service service.Servicer
}

func New(s service.Servicer) Handler {
	return handler{
		service: s,
	}
}
