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
	AsignarImagenPrincipal(idImagen *bson.ObjectID, ctx context.Context) error
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

func (r *imagen) AsignarImagenPrincipal(idImagen *bson.ObjectID, ctx context.Context) error {

	var imagen struct {
		producto bson.ObjectID `bson:"producto"`
	}

	err := r.collection.FindOne(
		ctx,
		bson.M{"_id": idImagen},
	).Decode(&imagen)

	if err != nil {
		return err
	}

	_, err = r.collection.UpdateMany(
		ctx,
		bson.M{
			"prodcuto": imagen.producto,
		},
		bson.M{
			"$set": bson.M{
				"principal": false,
			},
		},
	)

	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{
			"_id": idImagen,
		},
		bson.M{
			"$set": bson.M{
				"principal": true,
			},
		},
	)

	return err
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
