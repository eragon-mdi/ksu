package router

import (
	"net/http"

	"github.com/eragon-mdi/ksu/pkg/apperrors"
	applog "github.com/eragon-mdi/ksu/pkg/log"
	"github.com/labstack/echo/v4"
)

type Router interface {
	Echo() *echo.Echo
	Handler() http.Handler
}

type router struct {
	router *echo.Echo
}

func New() Router {
	e := echo.New()
	e.HTTPErrorHandler = apperrors.CustomHTTPErrorHandler
	e.Use(applog.InitMiddlewareLogging())

	return &router{
		router: e,
	}
}

func (r router) Echo() *echo.Echo {
	return r.router.AcquireContext().Echo()
}

func (r router) Handler() http.Handler {
	return r.router.AcquireContext().Echo()
}
