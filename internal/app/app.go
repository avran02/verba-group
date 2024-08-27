package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/avran02/verba-group/config"
	"github.com/avran02/verba-group/internal/controller"
	"github.com/avran02/verba-group/internal/repository"
	"github.com/avran02/verba-group/internal/router"
	"github.com/avran02/verba-group/internal/service"
	"github.com/avran02/verba-group/logger"
)

type App struct {
	router     *router.Router
	config     *config.Config
	repository repository.Repository
}

func (a *App) Run() error {
	serverEndpoint := fmt.Sprintf("%s:%s", a.config.Server.Host, a.config.Server.Port)
	slog.Info("Starting server at " + serverEndpoint)
	s := http.Server{
		Addr:    serverEndpoint,
		Handler: a.router,
	}

	s.RegisterOnShutdown(func() {
		if err := a.repository.Close(); err != nil {
			slog.Error("can't close db conn: " + err.Error())
		}
	})

	return s.ListenAndServe()
}

func New() *App {
	config := config.New()
	logger.Setup(config.Server)

	repository := repository.New(config.Postgres)
	service := service.New(repository)
	controller := controller.New(service)
	router := router.New(controller)

	return &App{
		router: router,
		config: config,
	}
}
