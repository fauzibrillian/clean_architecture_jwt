package barang

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Barang struct {
	ID          uint
	ProductName string
	Stock       uint
	Price       uint
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	TambahBarang(token *jwt.Token, newBarang Barang) (Barang, error)
}

type Repo interface {
	InsertBarang(userId uint, newBarang Barang) (Barang, error)
}
