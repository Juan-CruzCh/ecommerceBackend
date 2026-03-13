package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VarianteProducto struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Talla    string        `bson:"talla"`
	Color    string        `bson:"color"`
	Producto bson.ObjectID `bson:"producto"`
	Fecha    time.Time     `bson:"fecha"`
	flag     enum.FlagE    `bson:"flag"`
}
