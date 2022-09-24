package middleware

import (
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/internal/user"
)

type Manager struct {
	origins []string
	cfg     *config.Config
	userUC  user.UseCase
}

func NewMiddlewareManager(origins []string, cfg *config.Config, userUC user.UseCase) *Manager {
	return &Manager{origins: origins, cfg: cfg, userUC: userUC}
}
