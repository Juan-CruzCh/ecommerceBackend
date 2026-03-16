package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type DataStockDto struct {
	Stock []StockDto `json:"stock" validate:"required,min=1,dive"`
}

type StockDto struct {
	Cantidad int           `json:"cantidad" validate:"required,gt=0"`
	Talla    bson.ObjectID `json:"talla" validate:"required"`
	Producto bson.ObjectID `json:"producto" validate:"required"`
}
