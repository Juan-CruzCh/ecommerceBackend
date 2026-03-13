package repository

import (
	"context"
	"ecommerceBackend/src/module/producto/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VarianteProducto interface {
	CrearVarianteProducto(ctx context.Context, variante *model.VarianteProducto) (*mongo.InsertOneResult, error)
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

func (r *varianteProducto) CrearVarianteProducto(ctx context.Context, variante *model.VarianteProducto) (*mongo.InsertOneResult, error) {
	resultado, err := r.collection.InsertOne(ctx, variante)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
