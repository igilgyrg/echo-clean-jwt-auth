package usecase

import (
	"context"
	"github.com/igilgyrg/todo-echo/internal/domain"
	"github.com/igilgyrg/todo-echo/internal/user"
)

type userUC struct {
	repository user.Repository
}

func NewUserUC(repository user.Repository) user.UseCase {
	return &userUC{repository: repository}
}

func (u userUC) Get(ctx context.Context, ID domain.ID) (*domain.User, error) {
	return u.repository.Get(ctx, ID)
}

func (u userUC) Store(ctx context.Context, task *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userUC) Update(ctx context.Context, task *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userUC) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.repository.GetByEmail(ctx, email)
}
