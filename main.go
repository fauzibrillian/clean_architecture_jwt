package main

import (
	"clean_architecture_jwt/config"
	"clean_architecture_jwt/features/barang"
	"clean_architecture_jwt/features/user"
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
	db.AutoMigrate(&model.UserModel{}, &model.ProductModel{})
	m := model.UserQuery{DB: db}
	userController := user.UserController{Model: m}

	bm := model.BarangQuery{DB: db}
	barangController := barang.BarangController{Model: bm}

	routes.InitRoute(e, userController, barangController)

	e.Logger.Fatal(e.Start(":8000"))
}
