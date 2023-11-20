package handler

type BarangRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
}

type BarangResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Stock       uint   `json:"stock"`
	Price       uint   `json:"price"`
}
