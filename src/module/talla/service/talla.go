package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/talla/dto"
	"ecommerceBackend/src/module/talla/model"
	"ecommerceBackend/src/module/talla/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Talla struct {
	tallaRepository repository.Talla
}

func NewTallaService(repo repository.Talla) Talla {
	return Talla{
		tallaRepository: repo,
	}
}

func (s *Talla) ListarTallas(ctx context.Context) (*[]model.Talla, error) {
	data, err := s.tallaRepository.ListarTalla(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Talla) CrearTallas(body *dto.TallaDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	data := model.Talla{
		Nombre: body.Nombre,
		Flag:   enum.FlagNuevo,
		Fecha:  utils.FechaHoraBolivia(),
	}
	resultado, err := s.tallaRepository.CrearTalla(&data, ctx)
	if err != nil {
		return nil, err
	}

	return resultado, nil
}
