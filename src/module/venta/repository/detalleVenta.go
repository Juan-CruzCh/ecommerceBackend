package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetalleVenta interface {
	CrearDetalleVenta(ctx context.Context)
	ListarDetalleVenta(ctx context.Context)
}

type detalleVenta struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewDetalleVentaRepository(db *mongo.Database) DetalleVenta {
	return &detalleVenta{
		db:         db,
		collection: db.Collection("DetalleVenta"),
	}
}

func (r *detalleVenta) CrearDetalleVenta(ctx context.Context) {

}

func (r *detalleVenta) ListarDetalleVenta(ctx context.Context) {

}
