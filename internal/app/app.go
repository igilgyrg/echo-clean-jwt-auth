package app

import (
	"context"
	"fmt"
	"github.com/igilgyrg/todo-echo/internal/config"
	"github.com/igilgyrg/todo-echo/pkg/logging"
	mongoclient "github.com/igilgyrg/todo-echo/pkg/repository/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	Timeout        = 10
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type App struct {
	cfg     *config.Config
	mongoDB *mongo.Database
	echo    *echo.Echo
}

func NewApp() *App {
	cfg := config.NewConfig()
	mongoConfig := mongoclient.NewMongoConfig(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase, cfg.MongoUsername, cfg.MongoPassword)

	mongoClient, err := mongoclient.Init(mongoConfig)
	if err != nil {
		log.Fatal("mongo database have not inited")
	}

	return &App{
		cfg:     cfg,
		mongoDB: mongoClient,
		echo:    echo.New(),
	}
}

func (a *App) Start() error {
	httpServer := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%s", a.cfg.Port),
		ReadTimeout:    Timeout * time.Second,
		WriteTimeout:   Timeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		if err := a.echo.StartServer(httpServer); err != nil {
			log.Fatalf("error start server, %w", err)
		}
	}()

	if err := a.MapHandlers(a.echo, a.cfg); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(logging.ContextWithLogger(context.Background()), ctxTimeout*time.Second)
	defer shutdown()

	return a.echo.Shutdown(ctx)
}
