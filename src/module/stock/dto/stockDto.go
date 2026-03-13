package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type StockDto struct {
	Cantidad         int           `json:"cantidad" validate:"required,gt=0"`
	Precio           int           `json:"precio" validate:"required,gt=0"`
	VarianteProducto bson.ObjectID `json:"varianteProducto" validate:"required"`
	Producto         bson.ObjectID `json:"producto" validate:"required"`
}
