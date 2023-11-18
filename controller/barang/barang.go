package barang

import (
	"clean_architecture_jwt/model"
	"clean_architecture_jwt/utils/jwt"
	"net/http"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BarangController struct {
	Model model.BarangQuery
}

func (bc *BarangController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := jwt.ExtractToken(c.Get("user").(*gojwt.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "tidak ada kuasa untuk mengakses",
			})
		}
		var input = new(BarangRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input tidak sesuai",
			})
		}
		var inputProses = new(model.ProductModel)
		inputProses.ProductName = input.ProductName
		inputProses.Stock = uint(input.Stock)
		inputProses.Price = uint(input.Price)
		inputProses.UserID = userid

		result, err := bc.Model.AddBarang(*inputProses)
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "dobel input nama",
				})
			}
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input tidak sesuai",
			})
		}
		var response = new(BarangResponse)
		response.ID = result.ID
		response.ProductName = result.ProductName
		response.Stock = uint(result.Stock)
		response.Price = uint(result.Price)
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}
