package repository

import (
	"context"
	"ecommerceBackend/src/module/producto/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Imagen interface {
	CrearImgen(ctx context.Context, imagen *model.Imagen)
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
