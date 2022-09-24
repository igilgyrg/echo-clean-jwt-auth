package http

import (
	"github.com/igilgyrg/todo-echo/internal/auth"
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/internal/domain"
	error2 "github.com/igilgyrg/todo-echo/internal/error"
	"github.com/igilgyrg/todo-echo/internal/user"
	"github.com/igilgyrg/todo-echo/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authHandler struct {
	cfg         *config.Config
	usecase     auth.UseCase
	userUsecase user.UseCase
}

func NewAuthHandler(cfg *config.Config, usecase auth.UseCase, userUsecase user.UseCase) auth.Handler {
	return &authHandler{cfg: cfg, usecase: usecase, userUsecase: userUsecase}
}

func (a authHandler) Login() echo.HandlerFunc {
	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c echo.Context) error {
		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			return error2.NewBadRequestError(err)
		}

		token, err := a.usecase.Login(c.Request().Context(), &domain.User{
			Email:    login.Email,
			Password: login.Password,
		})

		if err != nil {
			return error2.NewUnauthorizedError(err)
		}

		return c.JSON(http.StatusOK, token)
	}
}

func (a authHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{}{})
	}
}

func (a authHandler) Register() echo.HandlerFunc {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c echo.Context) error {
		user := &User{}
		if err := utils.ReadRequest(c, user); err != nil {
			return error2.NewBadRequestError(err)
		}

		userByEmail, err := a.userUsecase.GetByEmail(c.Request().Context(), user.Email)
		if err != nil {
			return error2.NewSystemError(err)
		}

		if userByEmail != nil {
			return error2.NewUserIsExistsWithEmail(err)
		}

		token, err := a.usecase.Register(c.Request().Context(), &domain.User{
			ID:        domain.NewID(),
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: time.Now(),
		})

		if err != nil {
			return error2.NewSystemError(err)
		}

		return c.JSON(http.StatusOK, token)
	}
}

func (a authHandler) Refresh() echo.HandlerFunc {
	type Refresh struct {
		RefreshToken string `json:"refresh_token"`
	}
	return func(c echo.Context) error {
		refresh := &Refresh{}
		if err := utils.ReadRequest(c, refresh); err != nil {
			return error2.NewBadRequestError(err)
		}

		token, err := a.usecase.Refresh(c.Request().Context(), refresh.RefreshToken)
		if err != nil {
			return error2.NewUnauthorizedError(err)
		}

		return c.JSON(http.StatusOK, token)
	}
}
