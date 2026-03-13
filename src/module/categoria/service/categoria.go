package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/categoria/dto"
	"ecommerceBackend/src/module/categoria/model"
	"ecommerceBackend/src/module/categoria/repository"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Categoria struct {
	categoriaRepository repository.Categoria
}

func NewcategoriaService(repo repository.Categoria) Categoria {
	return Categoria{
		categoriaRepository: repo,
	}
}

func (s *Categoria) ListarCategoria(ctx context.Context) (*[]model.Categoria, error) {
	resultado, err := s.categoriaRepository.ListarCategoria(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *Categoria) CrearCategoria(ctx context.Context, body dto.CategoriaDto) (map[string]string, error) {
	var data model.Categoria = model.Categoria{
		Nombre: body.Nombre,
		Fecha:  utils.FechaHoraBolivia(),
		Flag:   enum.FlagNuevo,
	}
	resultado, err := s.categoriaRepository.CrearCategoria(ctx, &data)

	if err != nil {
		return nil, err
	}
	id, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("Error de parseo")
	}
	r := map[string]string{
		"categoria": id.Hex(),
	}
	return r, nil
}
