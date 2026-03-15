package model

import (
	"ecommerceBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Producto struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Codigo          string        `bson:"codigo"`
	Nombre          string        `bson:"nombre"`
	Descripcion     string        `bson:"descripcion"`
	Categoria       bson.ObjectID `bson:"categoria"`
	Fecha           time.Time     `bson:"fecha"`
	Flag            enum.FlagE    `bson:"flag"`
	Publico         bool          `bson:"publico"`
	Destacado       bool          `bson:"destacado"`
	PrecioCompra    float64       `bson:"precioCompra"`
	PrecioVenta     float64       `bson:"precioVenta"`
	ImagenPrincipal string        `bson:"imagenPrincipal"`
}
