package handler

import (
	"clean_architecture_jwt/features/barang"
	"net/http"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BarangHandler struct {
	s barang.Service
}

func New(s barang.Service) barang.Handler {
	return &BarangHandler{
		s: s,
	}

}

func (bc *BarangHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(BarangRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(barang.Barang)
		inputProcess.ProductName = input.ProductName
		inputProcess.Stock = input.Stock
		inputProcess.Price = input.Price

		result, err := bc.s.TambahBarang(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var response = new(BarangResponse)
		response.ProductName = result.ProductName
		response.Stock = result.Stock
		response.Price = result.Price
		response.ID = result.ID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}
