package main

import (
	"github.com/hostingvk4/badgerList/internal/handler"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/internal/repository/database"
	"github.com/hostingvk4/badgerList/internal/server"
	"github.com/hostingvk4/badgerList/internal/service"
	"log"
)

func main() {
	db := database.NewPostgresDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
