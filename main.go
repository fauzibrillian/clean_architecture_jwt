package main

import (
	"clean_architecture_jwt/config"
	"clean_architecture_jwt/controller/user"
	"clean_architecture_jwt/model"
	"clean_architecture_jwt/routes"
	"clean_architecture_jwt/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start server kesalahan database")
	}

	db, err := database.InitMySql(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}
	db.AutoMigrate(&model.UserModel{})
	model := model.UserQuery{DB: db}
	userController := user.UserController{Model: model}

	routes.InitRoute(e, userController)

	e.Logger.Fatal(e.Start(":8000"))
}
