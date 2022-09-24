package http

import (
	"github.com/igilgyrg/todo-echo/internal/auth"
	"github.com/igilgyrg/todo-echo/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroups *echo.Group, h auth.Handler, manager *middleware.Manager) {
	authGroups.POST("/login", h.Login(), manager.ResponseErrorToJSON())
	authGroups.POST("/register", h.Register(), manager.ResponseErrorToJSON())
	authGroups.POST("/logout", h.Logout(), manager.ResponseErrorToJSON(), manager.AuthJWTMiddleware())
	authGroups.POST("/refresh", h.Refresh(), manager.ResponseErrorToJSON())
}
