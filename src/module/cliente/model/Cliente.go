package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Cliente struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Codigo    string        `bson:"codigo" json:"codigo"`
	Nombre    string        `bson:"nombre" json:"nombre"`
	Apellidos string        `bson:"Apellidos" json:"Apellidos"`
	Flag      enum.FlagE    `bson:"flag" json:"flag"`
	Fecha     time.Time     `bson:"fecha"  json:"fecha"`
	Celular   string        `bson:"celular" json:"celular"`
}
