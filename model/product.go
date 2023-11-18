package model

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	ProductName string
	Stock       uint
	Price       uint
	UserID      uint
}
type BarangQuery struct {
	DB *gorm.DB
}

func (bq *BarangQuery) AddBarang(newBarang ProductModel) (ProductModel, error) {
	if err := bq.DB.Create(&newBarang).Error; err != nil {
		return ProductModel{}, err
	}

	return newBarang, nil
}
