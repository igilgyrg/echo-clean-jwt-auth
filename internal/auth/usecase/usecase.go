package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/igilgyrg/todo-echo/internal/auth"
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/internal/domain"
	internalerror "github.com/igilgyrg/todo-echo/internal/error"
	"github.com/igilgyrg/todo-echo/internal/user"
	"time"
)

const (
	AccessTokenExpiredMinutes = 15
	RefreshTokenExpiredHours  = 24
)

type AuthUC struct {
	repository user.Repository
	cfg        *config.Config
}

func NewAuthUseCase(repository user.Repository, cfg *config.Config) auth.UseCase {
	return &AuthUC{repository: repository, cfg: cfg}
}

func (a AuthUC) Login(ctx context.Context, user *domain.User) (*domain.Token, error) {
	u, err := a.repository.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err := u.ComparePassword(user.Password); err != nil {
		// TODO invalid password error
		return nil, err
	}

	accessToken, err := generateAccessToken(u.ID, a.cfg.AccessTokenSignature)
	if err != nil {
		return nil, errors.New("")
	}

	refreshToken, err := generateRefreshToken(u, a.cfg.RefreshTokenSignature)
	if err != nil {
		return nil, errors.New("")
	}

	return &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a AuthUC) Register(ctx context.Context, user *domain.User) (*domain.Token, error) {
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	id, err := a.repository.Store(ctx, user)
	user.ID = id
	if err != nil {
		return nil, err
	}

	accessToken, err := generateAccessToken(id, a.cfg.AccessTokenSignature)
	if err != nil {
		return nil, internalerror.NewAccessTokenInvalid(err)
	}

	refreshToken, err := generateRefreshToken(user, a.cfg.RefreshTokenSignature)
	if err != nil {
		return nil, internalerror.NewRefreshTokenInvalid(err)
	}

	return &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (a AuthUC) Refresh(ctx context.Context, refreshToken string) (*domain.Token, error) {
	userID, ok := validateRefreshToken(refreshToken, a.cfg.RefreshTokenSignature)
	if ok {
		accessToken, err := generateAccessToken(domain.ID(userID), a.cfg.AccessTokenSignature)
		if err != nil {
			return nil, internalerror.NewAccessTokenInvalid(err)
		}

		return &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
	}
	return nil, internalerror.NewRefreshTokenInvalid(nil)
}

func generateAccessToken(userID domain.ID, accessTokenSignature string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(AccessTokenExpiredMinutes * time.Minute).Unix()
	claims["auth"] = true
	claims["sub"] = userID

	accessToken, err := token.SignedString([]byte(accessTokenSignature))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func generateRefreshToken(u *domain.User, refreshTokenSignature string) (string, error) {
	rt := jwt.New(jwt.SigningMethodHS256)
	rtClaims := rt.Claims.(jwt.MapClaims)
	rtClaims["sub"] = u.ID
	rtClaims["exp"] = time.Now().Add(RefreshTokenExpiredHours * time.Hour).Unix()

	refreshToken, err := rt.SignedString([]byte(refreshTokenSignature))
	if err != nil {
		return "", err
	}

	return refreshToken, nil

}

func validateRefreshToken(refreshTokenString string, refreshTokenSignature string) (string, bool) {
	if refreshTokenString == "" {
		return "", false
	}

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(refreshTokenSignature)
		return secret, nil
	})

	if err != nil {
		return "", false
	}

	if !refreshToken.Valid {
		return "", false
	}

	userID := ""

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		userID, ok = claims["sub"].(string)
		if !ok {
			return "", false
		}
	}

	return userID, true
}
