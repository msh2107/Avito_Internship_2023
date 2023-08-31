package app

import (
	"Avito/config"
	v1 "Avito/internal/controller/http/v1"
	"Avito/internal/repository"
	"Avito/internal/service"
	"Avito/pkg/db/postgresql"
	"Avito/pkg/httpserver"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	SetLogrus("debug")

	log.Info("Initializing postgres...")
	pg, err := postgresql.New(cfg.PG.URL, postgresql.MaxPoolSize(cfg.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}
	defer pg.Close()

	log.Info("Initializing repositories...")
	repos := repository.NewRepositories(pg)

	log.Info("Initializing services...")
	services := service.NewServices(repos)

	log.Info("Initializing handlers and routes...")
	handler := gin.New()
	v1.NewRouter(handler, services)

	log.Info("Starting http server...")
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
