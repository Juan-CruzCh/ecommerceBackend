package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type VarianteProductoDto struct {
	Talla    string        `json:"talla" validate:"required"`
	Color    string        `json:"color" validate:"required"`
	Producto bson.ObjectID `json:"producto" validate:"required"`
}
