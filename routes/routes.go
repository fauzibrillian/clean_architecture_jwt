package routes

import (
	"clean_architecture_jwt/features/barang"
	"clean_architecture_jwt/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.Handler, bc barang.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
	RouteBarang(e, bc)
}

func RouteUser(e *echo.Echo, uc user.Handler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}

func RouteBarang(e *echo.Echo, bc barang.Handler) {
	e.POST("/barangs", bc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
