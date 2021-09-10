package main

import (
	user2 "github.com/dico87/users/api/handlers/user"
	"github.com/dico87/users/api/routes"
	"github.com/dico87/users/model"
	"github.com/dico87/users/repository"
	"github.com/dico87/users/service"
	echo2 "github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const host = "127.0.0.1"
const port = "3306"
const dbName = "users"
const user = "root"
const password = "12345"

func main() {
	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connected error : %s", err.Error())
		os.Exit(3)
	} else {
		db.AutoMigrate(&model.DocumentType{}, &model.User{})

		userRepository := repository.NewMySqlRepository(db)
		userService := service.NewUserService(userRepository)

		echo := echo2.New()
		handler := user2.NewUserHandler(userService)

		routes.InitRoutes(echo, handler)

		err := echo.Start(":9090")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
