package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/internal/domain"
	internalerror "github.com/igilgyrg/todo-echo/internal/error"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (m *Manager) AuthJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			if bearerHeader != "" {
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					// TODO logger
					return internalerror.NewUnauthorizedError(internalerror.Unauthorized)
				}

				tokenString := headerParts[1]

				err := m.validateJWTToken(tokenString, c, m.cfg)
				if err != nil {
					return internalerror.NewUnauthorizedError(err)
				}

				return next(c)
			}
			return nil
		}
	}
}

func (m *Manager) validateJWTToken(tokenString string, c echo.Context, cfg *config.Config) error {
	if tokenString == "" {
		return errors.New("")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(cfg.AccessTokenSignature)
		return secret, nil
	})

	if err != nil {
		return errors.New("")
	}

	if !token.Valid {
		return errors.New("")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return errors.New("")
		}

		user, err := m.userUC.Get(c.Request().Context(), domain.ID(userID))
		if err != nil {
			return errors.New("")
		}

		c.Set("user", user)

		ctx := context.WithValue(c.Request().Context(), "user", user)
		c.SetRequest(c.Request().WithContext(ctx))
	}

	return nil
}

func (m Manager) ResponseErrorToJSON() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "application/json")
			var appError *internalerror.HttpError
			err := next(c)
			if err != nil {
				if errors.As(err, &appError) {
					c.Response().WriteHeader(appError.Status())
					c.Response().Write(appError.Marshal())
					return nil
				}

				c.Response().WriteHeader(http.StatusInternalServerError)
				errBytes, _ := json.Marshal(err)
				c.Response().Write(errBytes)
			}
			c.Response().WriteHeader(http.StatusOK)
			return nil
		}
	}

}
