package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VarianteProducto interface {
	CrearVarianteProducto(ctx context.Context)
}

type varianteProducto struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewVarianteProductoRepository(db *mongo.Database) VarianteProducto {
	return &varianteProducto{
		db:         db,
		collection: db.Collection("VarianteProducto"),
	}
}

func (r *varianteProducto) CrearVarianteProducto(ctx context.Context) {

}
