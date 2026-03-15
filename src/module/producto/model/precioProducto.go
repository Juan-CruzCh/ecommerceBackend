package model

import "go.mongodb.org/mongo-driver/v2/bson"

type PrecioProducto struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Producto     bson.ObjectID `bson:"producto"`
	PrecioCompra float64       `bson:"precioCompra"`
	PrecioVenta  float64       `bson:"precioVenta"`
}
