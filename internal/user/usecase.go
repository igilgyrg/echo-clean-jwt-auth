package user

import (
	"context"
	"github.com/igilgyrg/todo-echo/internal/domain"
)

type UseCase interface {
	Get(ctx context.Context, ID domain.ID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Store(ctx context.Context, task *domain.User) error
	Update(ctx context.Context, task *domain.User) error
}
