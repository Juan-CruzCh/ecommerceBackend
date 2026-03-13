package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/producto/model"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Producto interface {
	CrearProducto(ctx context.Context, producto *model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(ctx context.Context)
	EditarProducto(ctx context.Context)
	ListarProducto(ctx context.Context)
}

type producto struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProductoRepository(db *mongo.Database) Producto {
	return &producto{
		db:         db,
		collection: db.Collection("Producto"),
	}
}

func (r *producto) CrearProducto(ctx context.Context, producto *model.Producto) (*mongo.InsertOneResult, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"nombre": producto.Nombre, "flag": enum.FlagNuevo})
	if cantidad > 0 {
		return nil, errors.New("El producto ya existe")
	}
	resultado, err := r.collection.InsertOne(ctx, producto)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *producto) EliminarProducto(ctx context.Context) {

}

func (r *producto) EditarProducto(ctx context.Context) {

}

func (r *producto) ListarProducto(ctx context.Context) {

}
