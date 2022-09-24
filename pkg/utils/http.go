package utils

import (
	"errors"
	"github.com/igilgyrg/todo-echo/internal/domain"
	"github.com/labstack/echo/v4"
)

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return ValidateStruct(ctx.Request().Context(), request)
}

func UserFromContext(ctx echo.Context) (*domain.User, error) {
	user := ctx.Get("user").(*domain.User)
	if user == nil {
		return nil, errors.New("user have not founded from context")
	}
	return user, nil
}
