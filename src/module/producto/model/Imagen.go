package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Imagen struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	VarianteProducto bson.ObjectID `bson:"varianteProducto" json:"varianteProducto"`
	Path             string        `bson:"path" json:"path"`
	Fecha            time.Time     `bson:"fecha"`
	flag             enum.FlagE    `bson:"flag"`
}
