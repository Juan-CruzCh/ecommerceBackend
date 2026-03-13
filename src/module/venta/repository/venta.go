package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Venta interface {
	CrearVenta(ctx context.Context)
	ListarVenta(ctx context.Context)
}

type venta struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewVentaRepository(db *mongo.Database) Venta {
	return &venta{
		db:         db,
		collection: db.Collection("Venta"),
	}
}

func (r *venta) CrearVenta(ctx context.Context) {

}

func (r *venta) ListarVenta(ctx context.Context) {

}
