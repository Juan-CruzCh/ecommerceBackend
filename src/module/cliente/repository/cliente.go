package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type Cliente interface {
}
type cliente struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewClienteRepository(db *mongo.Database) Cliente {
	return &cliente{
		db:         db,
		collection: db.Collection("Cliente"),
	}
}
