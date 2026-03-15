package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/producto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Imagen interface {
	CrearImgen(ctx context.Context, imagen *model.Imagen)
	ListarImagenes(variante *bson.ObjectID, ctx context.Context) (*[]model.Imagen, error)
}

type imagen struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewImagenRepository(db *mongo.Database) Imagen {
	return &imagen{
		db:         db,
		collection: db.Collection("Imagen"),
	}
}

func (r *imagen) CrearImgen(ctx context.Context, imagen *model.Imagen) {
	r.collection.InsertOne(ctx, imagen)
}

func (r *imagen) ListarImagenes(variante *bson.ObjectID, ctx context.Context) (*[]model.Imagen, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo, "varianteProducto": variante})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.Imagen = []model.Imagen{}
	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}
	return &data, nil
}
