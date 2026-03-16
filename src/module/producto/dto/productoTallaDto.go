package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type ProductoTallaDto struct {
	Talla    bson.ObjectID `json:"talla" validate:"required"`
	Producto bson.ObjectID `json:"producto" validate:"required"`
}
