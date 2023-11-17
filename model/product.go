package model

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	ProductName string
	Stock       int
	Qty         int
	Price       int
	UserID      uint
}
