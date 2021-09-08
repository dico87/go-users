package main

import (
	"github.com/dico87/users/model"
	"github.com/dico87/users/repository"
	"github.com/dico87/users/service"
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
	url := user + ":" + password + "@tcp("+host+":"+port+")/"+dbName+"?charset=utf8&parseTime=True"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connected error : %s", err.Error())
		os.Exit(3)
	} else {
		db.AutoMigrate(&model.DocumentType{}, &model.User{})

		userRepository := repository.NewMySqlRepository(db)
		userService := service.NewUserService(userRepository)

		//user := model.User{
		//	DocumentTypeID: 1,
		//	Document: "1018408447",
		//	LastName: "Cortés",
		//	SurName: "Plazas",
		//	Name: "Diego",
		//	OtherNames: "Fernando",
		//	Birthday: time.Date(1987, 1, 19, 0,0,0,0, time.UTC),
		//	Sex: "M",
		//	Active: true,
		//}

		user, err := userService.FindByDocument(1, "101840447")

		if err != nil {
			log.Fatal(err)
		}

		log.Println(user)
	}
}
