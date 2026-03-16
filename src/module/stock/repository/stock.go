package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/stock/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Stock interface {
	CrearStock(stock *model.Stock, tx context.Context) error
	VerificarStock(producto *bson.ObjectID, ProductoTalla *bson.ObjectID, ctx context.Context) (*model.Stock, error)
	ActualizarStock(stock *bson.ObjectID, cantidad int, ctx context.Context) *mongo.SingleResult
}

type stock struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStockRepository(db *mongo.Database) Stock {
	return &stock{
		db:         db,
		collection: db.Collection("Stock"),
	}
}

func (r *stock) CrearStock(stock *model.Stock, ctx context.Context) error {
	_, err := r.collection.InsertOne(ctx, stock)
	if err != nil {
		return err
	}
	return nil

}
func (r *stock) ActualizarStock(stock *bson.ObjectID, cantidad int, ctx context.Context) *mongo.SingleResult {
	resultado := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": stock, "flag": enum.FlagNuevo}, bson.D{{Key: "$set", Value: bson.D{{Key: "cantidad", Value: cantidad}}}})
	return resultado
}

func (r *stock) VerificarStock(producto *bson.ObjectID, ProductoTalla *bson.ObjectID, ctx context.Context) (*model.Stock, error) {

	var stock model.Stock = model.Stock{}
	err := r.collection.FindOne(ctx, bson.M{"producto": producto, "productoTalla": ProductoTalla, "flag": enum.FlagNuevo}).Decode(&stock)
	if err != nil {
		return nil, err
	}
	return &stock, nil

}
