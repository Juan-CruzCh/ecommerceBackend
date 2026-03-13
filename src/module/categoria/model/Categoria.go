package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Categoria struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre string        `bson:"nombre" json:"nombre"`
	Fecha  time.Time     `bson:"fecha" json:"fecha"`
	Flag   enum.FlagE    `bson:"flag" json:"flag"`
}
