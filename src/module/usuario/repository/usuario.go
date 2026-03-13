package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Usuario interface {
	CrearUsuario(ctx context.Context)
	EliminarUsuario(ctx context.Context)
	EditarUsuario(ctx context.Context)
	ListarUsuario(ctx context.Context)
}

type usuario struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) Usuario {
	return &usuario{
		db:         db,
		collection: db.Collection("Usuario"),
	}
}

func (r *usuario) CrearUsuario(ctx context.Context) {

}

func (r *usuario) EliminarUsuario(ctx context.Context) {

}

func (r *usuario) EditarUsuario(ctx context.Context) {

}

func (r *usuario) ListarUsuario(ctx context.Context) {

}
