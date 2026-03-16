package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/talla/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Talla interface {
	CrearTalla(talla *model.Talla, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarTalla(ctx context.Context) (*[]model.Talla, error)
	EditarTalla(id *bson.ObjectID, nombre string, ctx context.Context) (*mongo.UpdateResult, error)
	EliminarTalla(id *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type talla struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTallaRepository(db *mongo.Database) Talla {
	return &talla{
		db:         db,
		collection: db.Collection("Talla"),
	}
}

func (r *talla) CrearTalla(talla *model.Talla, ctx context.Context) (*mongo.InsertOneResult, error) {

	resultado, err := r.collection.InsertOne(ctx, talla)
	if err != nil {
		return nil, err
	}

	return resultado, nil
}

func (r *talla) ListarTalla(ctx context.Context) (*[]model.Talla, error) {

	cursor, err := r.collection.Find(ctx, bson.M{
		"flag": enum.FlagNuevo,
	})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var data []model.Talla

	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *talla) EditarTalla(id *bson.ObjectID, nombre string, ctx context.Context) (*mongo.UpdateResult, error) {

	resultado, err := r.collection.UpdateOne(
		ctx,
		bson.M{
			"_id":  id,
			"flag": enum.FlagNuevo,
		},
		bson.M{
			"$set": bson.M{
				"nombre": nombre,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return resultado, nil
}

func (r *talla) EliminarTalla(id *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {

	resultado, err := r.collection.UpdateOne(
		ctx,
		bson.M{
			"_id": id,
		},
		bson.M{
			"$set": bson.M{
				"flag": enum.FlagEliminado,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return resultado, nil
}
