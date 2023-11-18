package main

import (
	"clean_architecture_jwt/config"
	"clean_architecture_jwt/routes"
	"clean_architecture_jwt/utils/database"

	bh "clean_architecture_jwt/features/barang/handler"
	br "clean_architecture_jwt/features/barang/repository"
	bs "clean_architecture_jwt/features/barang/service"
	uh "clean_architecture_jwt/features/user/handler"
	ur "clean_architecture_jwt/features/user/repository"
	us "clean_architecture_jwt/features/user/service"

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
	db.AutoMigrate(&ur.UserModel{}, &br.ProductModel{})

	userRepo := ur.New(db)
	userService := us.New(userRepo)
	userHandler := uh.New(userService)

	barangRepo := br.New(db)
	barangService := bs.New(barangRepo)
	barangHandler := bh.New(barangService)

	routes.InitRoute(e, userHandler, barangHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
