package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Imagen struct {
	ID               bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Producto         bson.ObjectID `bson:"producto" json:"producto"`
	VarianteProducto bson.ObjectID `bson:"varianteProducto" json:"varianteProducto"`
	Path             string        `bson:"path" json:"path"`
	Fecha            time.Time     `bson:"fecha" json:"fecha"`
	Flag             enum.FlagE    `bson:"flag" json:"flag"`
	Principal        bool          `bson:"principal" json:"principal"`
}
