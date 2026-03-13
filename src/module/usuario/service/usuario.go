package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/usuario/dto"
	"ecommerceBackend/src/module/usuario/model"
	"ecommerceBackend/src/module/usuario/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Usuario struct {
	usuarioRepository repository.Usuario
}

func NewUsuarioService(repo repository.Usuario) Usuario {
	return Usuario{
		usuarioRepository: repo,
	}
}
func (s *Usuario) CrearUsuario(body *dto.UsuarioDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	hash, err := utils.EncriptarPassword(*body.Password)
	var data model.Usuario = model.Usuario{
		Ci:        body.Ci,
		Nombre:    body.Nombre,
		Celular:   body.Celular,
		Apellidos: body.Apellidos,

		Usuario:  body.Usuario,
		Password: hash,

		Flag:  enum.FlagNuevo,
		Rol:   body.Rol,
		Fecha: utils.FechaHoraBolivia(),
	}
	resultado, err := s.usuarioRepository.CrearUsuario(&data, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
func (s *Usuario) ListarUsuarios(ctx context.Context) (*[]model.Usuario, error) {

	resultado, err := s.usuarioRepository.ListarUsuario(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
func (s *Usuario) Eliminar(id *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	resultado, err := s.usuarioRepository.EliminarUsuario(id, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (s *Usuario) ActualizarUsuario(id *bson.ObjectID, body *dto.UsuarioDto, ctx context.Context) (*mongo.UpdateResult, error) {

	var data model.Usuario = model.Usuario{
		Ci:        body.Ci,
		Nombre:    body.Nombre,
		Celular:   body.Celular,
		Apellidos: body.Apellidos,
		Usuario:   body.Usuario,

		Rol: body.Rol,
	}
	resultado, err := s.usuarioRepository.EditarUsuario(id, &data, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
