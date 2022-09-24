package auth

import (
	"context"
	"github.com/igilgyrg/todo-echo/internal/domain"
)

type UseCase interface {
	Login(ctx context.Context, user *domain.User) (*domain.Token, error)
	Register(ctx context.Context, user *domain.User) (*domain.Token, error)
	Refresh(ctx context.Context, refreshToken string) (*domain.Token, error)
}
