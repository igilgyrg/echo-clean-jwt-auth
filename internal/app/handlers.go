package app

import (
	authHttp "github.com/igilgyrg/todo-echo/internal/auth/delivery/http"
	authUseCase "github.com/igilgyrg/todo-echo/internal/auth/usecase"
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/internal/middleware"
	taskHttp "github.com/igilgyrg/todo-echo/internal/task/delivery/http"
	taskRepository "github.com/igilgyrg/todo-echo/internal/task/repository"
	taskUsecase "github.com/igilgyrg/todo-echo/internal/task/usecase"
	userRepository "github.com/igilgyrg/todo-echo/internal/user/repository"
	userUsecase "github.com/igilgyrg/todo-echo/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *App) MapHandlers(e *echo.Echo, cfg *config.Config) error {
	// Init repositories
	userRepo := userRepository.NewUserMongoRepository(a.mongoDB)
	taskRepo := taskRepository.NewTaskMongoRepository(a.mongoDB)

	//Init usecases
	authCase := authUseCase.NewAuthUseCase(userRepo, cfg)
	taskCase := taskUsecase.NewTaskUC(taskRepo)
	userCase := userUsecase.NewUserUC(userRepo)

	// Init handlers
	authHandlers := authHttp.NewAuthHandler(a.cfg, authCase, userCase)
	tasksHandlers := taskHttp.NewTasksHandler(a.cfg, taskCase)

	mw := middleware.NewMiddlewareManager([]string{"*"}, a.cfg, userCase)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	tasksGroup := v1.Group("/tasks")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	taskHttp.MapTasksRoutes(tasksGroup, tasksHandlers, mw)

	return nil
}
