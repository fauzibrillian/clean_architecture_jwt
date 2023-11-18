package routes

import (
	"clean_architecture_jwt/controller/barang"
	"clean_architecture_jwt/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.UserController, bc barang.BarangController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
	RouteBarang(e, bc)
}

func RouteUser(e *echo.Echo, uc user.UserController) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/users", uc.GetListUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func RouteBarang(e *echo.Echo, bc barang.BarangController) {
	e.POST("/barangs", bc.Register(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
