package routes

import (
	"spos/auth/controllers"

	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/middleware"
)

type route struct {
	middleware middleware.Clients
	controller controllers.Controller
}

type Route interface {
	// Router is base router
	Router() *echo.Echo
}

func NewRouter(mdl middleware.Clients, ctrl controllers.Controller) Route {
	return &route{
		middleware: mdl,
		controller: ctrl,
	}
}

func (r *route) Router() *echo.Echo {
	router := echo.New()
	// base path
	auth := router.Group("", echo.WrapMiddleware(r.middleware.APIKey))
	auth.POST("/login", r.controller.LoginHandler)

	return router
}
