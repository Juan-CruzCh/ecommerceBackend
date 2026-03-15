package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/producto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VarianteProducto interface {
	CrearVarianteProducto(ctx context.Context, variante *model.VarianteProducto) (*mongo.InsertOneResult, error)
	ListarVarianteProducto(producto *bson.ObjectID, ctx context.Context) (*[]model.VarianteProducto, error)
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

func (r *varianteProducto) ListarVarianteProducto(producto *bson.ObjectID, ctx context.Context) (*[]model.VarianteProducto, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"producto": producto, "flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.VarianteProducto = []model.VarianteProducto{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}
