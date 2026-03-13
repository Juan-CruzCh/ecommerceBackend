package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Producto struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Nombre      string        `bson:"nombre"`
	Descripcion string        `bson:"Descripcion"`
	Categoria   bson.ObjectID `bson:"categoria"`
	Fecha       time.Time     `bson:"fecha"`
	flag        enum.FlagE    `bson:"flag"`
	Destacado   bool          `bson:"destacado"`
}
