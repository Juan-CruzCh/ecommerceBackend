package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Stock interface {
	CrearStock(ctx context.Context)
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

func (r *stock) CrearStock(ctx context.Context) {

}
