package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductoTalla interface {
	CrearProductoTalla(productoID, tallaID bson.ObjectID, ctx context.Context) (*mongo.InsertOneResult, error)
	BuscarProductoTalla(productoID, tallaID bson.ObjectID, ctx context.Context) (*model.ProductoTalla, error)
}

type productoTalla struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProductoTallaRepository(db *mongo.Database) ProductoTalla {
	return &productoTalla{
		db:         db,
		collection: db.Collection("ProductoTalla"),
	}
}
func (r *productoTalla) CrearProductoTalla(productoID, tallaID bson.ObjectID, ctx context.Context) (*mongo.InsertOneResult, error) {
	var productoTalla model.ProductoTalla = model.ProductoTalla{
		Talla:    tallaID,
		Producto: productoID,
		Fecha:    utils.FechaHoraBolivia(),
		Flag:     enum.FlagNuevo,
	}
	resultado, err := r.collection.InsertOne(ctx, productoTalla)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *productoTalla) BuscarProductoTalla(productoID, tallaID bson.ObjectID, ctx context.Context) (*model.ProductoTalla, error) {
	var productoTalla model.ProductoTalla = model.ProductoTalla{}
	err := r.collection.FindOne(ctx, bson.M{"producto": productoID, "talla": tallaID, "flag": enum.FlagNuevo}).Decode(&productoTalla)
	if err != nil {
		return nil, err
	}

	return &productoTalla, nil

}
