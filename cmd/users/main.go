package main

import (
	"github.com/dico87/users/internal/users/application"
	"github.com/dico87/users/internal/users/domain"
	"github.com/dico87/users/internal/users/infrastructure"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const host = "127.0.0.1"
const port = "3306"
const dbName = "users"
const user = "root"
const password = "12345"

func main() {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logger.Info("Application Start")


	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		logger.Fatalf("Database connected error : %s", err.Error())
		os.Exit(3)
	} else {
		db.AutoMigrate(&domain.DocumentType{}, &domain.User{})

		userRepository := infrastructure.NewMySqlRepository(db)
		userService := application.New(userRepository)

		echo := echo.New()
		handler := infrastructure.New(userService)

		infrastructure.InitRoutes(echo, handler)

		err := echo.Start(":9090")

		if err != nil {
			logger.Fatal(err.Error())
			os.Exit(3)
		}

		logger.Info("Http server start")
	}
}
