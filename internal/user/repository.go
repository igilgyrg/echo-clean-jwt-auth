package user

import (
	"context"
	"github.com/igilgyrg/todo-echo/internal/domain"
)

type Repository interface {
	Get(ctx context.Context, ID domain.ID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Store(ctx context.Context, user *domain.User) (domain.ID, error)
	Update(ctx context.Context, user *domain.User) error
}
