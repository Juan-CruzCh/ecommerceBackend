package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/model"
	"errors"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Producto interface {
	CrearProducto(ctx context.Context, producto *model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(ctx context.Context)
	EditarProducto(ctx context.Context)
	ListarProducto(ctx context.Context) (*[]bson.M, error)
	countDocumentsProducto(ctx context.Context) (int, error)
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
	countDocumentsProducto, err := r.countDocumentsProducto(ctx)
	if err != nil {
		return nil, err
	}

	producto.Codigo = "PROD-" + strconv.Itoa(countDocumentsProducto)
	resultado, err := r.collection.InsertOne(ctx, producto)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
func (r *producto) countDocumentsProducto(ctx context.Context) (int, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return int(cantidad + 1), nil

}

func (r *producto) EliminarProducto(ctx context.Context) {

}

func (r *producto) EditarProducto(ctx context.Context) {

}

func (r *producto) ListarProducto(ctx context.Context) (*[]bson.M, error) {

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "flag", Value: enum.FlagNuevo,
					},
				},
			},
		},
		utils.Lookup("Categoria", "categoria", "_id", "categoria"),

		bson.D{
			{
				Key: "$project", Value: bson.D{
					{
						Key: "_id", Value: 1,
					},
					{
						Key: "codigo", Value: 1,
					},
					{
						Key: "nombre", Value: 1,
					},
					{
						Key: "categoria", Value: utils.ArrayElemAt("$categoria.nombre", 0),
					},
					{
						Key: "publico", Value: 1,
					},
					{
						Key: "destacado", Value: 1,
					},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil
}
