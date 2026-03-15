package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Stock struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Cantidad         int           `bson:"cantidad"`
	VarianteProducto bson.ObjectID `bson:"varianteProducto"`
	Producto         bson.ObjectID `bson:"producto"`
	Fecha            time.Time     `bson:"fecha"`
	Flag             enum.FlagE    `bson:"flag"`
}
