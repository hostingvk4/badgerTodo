package main

import (
	"context"
	"github.com/hostingvk4/badgerList/internal/handler"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/internal/repository/database"
	"github.com/hostingvk4/badgerList/internal/server"
	"github.com/hostingvk4/badgerList/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title List App Api
// @version 1.0
// @description Api Server for List Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db := database.NewPostgresDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	println("app start")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	println("app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
