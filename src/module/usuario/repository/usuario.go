package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/module/usuario/model"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Usuario interface {
	CrearUsuario(usuario *model.Usuario, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarUsuario(ctx context.Context) (*[]model.Usuario, error)
	EliminarUsuario(usuario *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
	EditarUsuario(id *bson.ObjectID, usuario *model.Usuario, ctx context.Context) (*mongo.UpdateResult, error)
	/*BuscarUsuarioPorUsuario(usuario string, ctx context.Context) (*model.Usuario, error)
	BuscarUsuarioPorUsuarioId(usuario *bson.ObjectID, ctx context.Context) (*model.Usuario, error)*/
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

func (r *usuario) CrearUsuario(usuario *model.Usuario, ctx context.Context) (*mongo.InsertOneResult, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "usuario": usuario.Usuario})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, fmt.Errorf("El usuario ya existe")
	}
	resultado, err := r.collection.InsertOne(ctx, usuario)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *usuario) EliminarUsuario(usuario *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := r.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": usuario}, flagEliminado)
	if err != nil {
		return nil, err
	}
	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("El usuario no existe")
	}
	return resultado, nil
}

func (r *usuario) EditarUsuario(id *bson.ObjectID, usuario *model.Usuario, ctx context.Context) (*mongo.UpdateResult, error) {
	var filter bson.D = bson.D{
		{
			Key: "flag", Value: enum.FlagNuevo,
		},
		{
			Key: "usuario", Value: usuario.Usuario,
		},
		{
			Key: "_id", Value: bson.D{
				{Key: "$ne", Value: id},
			},
		},
	}

	cantidad, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, fmt.Errorf("El usuario ya existe")
	}
	var update bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{
				Key: "nombre", Value: usuario.Nombre,
			},
			{
				Key: "apellidos", Value: usuario.Apellidos,
			},

			{
				Key: "celular", Value: usuario.Celular,
			},
			{
				Key: "ci", Value: usuario.Ci,
			},

			{
				Key: "usuario", Value: usuario.Usuario,
			},
			{
				Key: "rol", Value: usuario.Rol,
			},
		}},
	}

	resultado, err := r.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": id}, update)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *usuario) ListarUsuario(ctx context.Context) (*[]model.Usuario, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.Usuario = []model.Usuario{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
