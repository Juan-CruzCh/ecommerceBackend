package service

import (
	"context"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/dto"
	"ecommerceBackend/src/module/producto/model"
	"ecommerceBackend/src/module/producto/repository"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Producto struct {
	productoRepository         repository.Producto
	varianteProductoRepository repository.VarianteProducto
}

func NewProductoService(productoRepository repository.Producto, varianteProductoRepository repository.VarianteProducto) Producto {
	return Producto{
		productoRepository:         productoRepository,
		varianteProductoRepository: varianteProductoRepository,
	}
}

func (s *Producto) CrearProducto(ctx context.Context, producto *dto.ProductoDto) (map[string]string, error) {

	var body model.Producto = model.Producto{
		Nombre:      producto.Nombre,
		Descripcion: producto.Descripcion,
		Categoria:   producto.Categoria,
		Destacado:   *producto.Destacado,
		Fecha:       utils.FechaHoraBolivia(),
	}
	resultado, err := s.productoRepository.CrearProducto(ctx, &body)
	if err != nil {
		return nil, err
	}
	id, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("Error de parseo")
	}
	data := map[string]string{
		"producto": id.Hex(),
	}

	return data, nil
}
