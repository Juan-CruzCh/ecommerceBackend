package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/categoria/model"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Categoria interface {
	CrearCategoria(ctx context.Context, categoria *model.Categoria) (*mongo.InsertOneResult, error)
	EliminarCategoria(ctx context.Context)
	EditarCategoria(ctx context.Context)
	ListarCategoria(ctx context.Context) (*[]model.Categoria, error)
}

type categoria struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewCategoriaRepository(db *mongo.Database) Categoria {
	return &categoria{
		db:         db,
		collection: db.Collection("Categoria"),
	}
}

func (r *categoria) CrearCategoria(ctx context.Context, categoria *model.Categoria) (*mongo.InsertOneResult, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"nombre": categoria.Nombre, "flag": enum.FlagNuevo})
	if cantidad > 0 {
		return nil, errors.New("La categoria ya existe")
	}
	resultado, err := r.collection.InsertOne(ctx, categoria)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *categoria) EliminarCategoria(ctx context.Context) {

}

func (r *categoria) EditarCategoria(ctx context.Context) {

}

func (r *categoria) ListarCategoria(ctx context.Context) (*[]model.Categoria, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	var resultado []model.Categoria = []model.Categoria{}
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return &resultado, nil
}
