package app

import (
	"context"
	"fmt"
	"github.com/hostingvk4/badgerList/internal/handler"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/internal/repository/database"
	"github.com/hostingvk4/badgerList/internal/server"
	"github.com/hostingvk4/badgerList/internal/service"
	"github.com/hostingvk4/badgerList/pkg/auth"
	"github.com/hostingvk4/badgerList/pkg/cipher"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing")
	}
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("db_pass"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
		return
	}
	cipherPass := cipher.NewCipher("aweawddsadfas23423asda")
	tokenAdministrator, err := auth.NewAdministrator("qweqsadawqe234324asdas")
	if err != nil {
		log.Fatalf("failed to initialize signing key : %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(service.ServicesConfig{
		Repos:              repos,
		TokenAdministrator: tokenAdministrator,
		RefreshTokenTTL:    24 * time.Hour,
		Cipher:             cipherPass,
	})
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
