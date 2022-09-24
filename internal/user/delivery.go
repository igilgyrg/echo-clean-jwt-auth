package user

import "github.com/labstack/echo/v4"

type Handler interface {
	Get() echo.HandlerFunc
	Save() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
