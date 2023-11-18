package repository

import (
	"clean_architecture_jwt/features/barang"

	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	ProductName string
	Stock       uint
	Price       uint
	UserID      uint
}
type BarangQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) barang.Repo {
	return &BarangQuery{
		db: db,
	}
}

func (bq *BarangQuery) InsertBarang(userID uint, newBarang barang.Barang) (barang.Barang, error) {
	var inputData = new(ProductModel)
	inputData.UserID = userID
	inputData.ProductName = newBarang.ProductName
	inputData.Stock = newBarang.Stock
	inputData.Price = newBarang.Price

	if err := bq.db.Create(&inputData).Error; err != nil {
		return barang.Barang{}, err
	}

	newBarang.ID = inputData.ID

	return newBarang, nil
}
