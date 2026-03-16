package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductoTalla struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Talla    bson.ObjectID `bson:"talla" json:"talla"`
	Producto bson.ObjectID `bson:"producto" json:"producto"`
	Fecha    time.Time     `bson:"fecha" json:"fecha"`
	Flag     enum.FlagE    `bson:"flag" json:"flag"`
}
