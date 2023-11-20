package service

import (
	"clean_architecture_jwt/features/barang"
	"clean_architecture_jwt/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type BarangService struct {
	m barang.Repo
}

func New(model barang.Repo) barang.Service {
	return &BarangService{
		m: model,
	}
}

func (bs *BarangService) TambahBarang(token *golangjwt.Token, newBarang barang.Barang) (barang.Barang, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return barang.Barang{}, err
	}
	result, err := bs.m.InsertBarang(userId, newBarang)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return barang.Barang{}, errors.New("dobel input")
		}
		return barang.Barang{}, errors.New("terjadi kesalahan")
	}
	return result, nil
}
