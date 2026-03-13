package service

import (
	"context"
	appUtils "ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/dto"
	"ecommerceBackend/src/module/producto/model"
	"ecommerceBackend/src/module/producto/repository"
	productoUtils "ecommerceBackend/src/module/producto/utils"
	"errors"
	"fmt"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Producto struct {
	productoRepository         repository.Producto
	varianteProductoRepository repository.VarianteProducto
	ImagenRepository           repository.Imagen
}

func NewProductoService(productoRepository repository.Producto, varianteProductoRepository repository.VarianteProducto, ImagenRepository repository.Imagen) Producto {
	return Producto{
		productoRepository:         productoRepository,
		varianteProductoRepository: varianteProductoRepository,
		ImagenRepository:           ImagenRepository,
	}
}

func (s *Producto) CrearProducto(ctx context.Context, producto *dto.ProductoDto) (map[string]string, error) {

	var body model.Producto = model.Producto{
		Nombre:      producto.Nombre,
		Descripcion: producto.Descripcion,
		Categoria:   producto.Categoria,
		Destacado:   *producto.Destacado,
		Fecha:       appUtils.FechaHoraBolivia(),
		Publico:     *producto.Publico,
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

func (s *Producto) CrearVarianteProducto(talla, color string, producto bson.ObjectID, imagenes []*multipart.FileHeader, ctx context.Context) (map[string]string, error) {
	variante := model.VarianteProducto{
		Talla:    talla,
		Color:    color,
		Producto: producto,
		Fecha:    appUtils.FechaHoraBolivia(),
	}
	resultado, err := s.varianteProductoRepository.CrearVarianteProducto(ctx, &variante)
	if err != nil {
		return nil, err
	}
	varianteId, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("Error de parseo")
	}
	for _, v := range imagenes {

		img, err := productoUtils.GuardarImagen(v)
		if err != nil {
			fmt.Println(err.Error())
		}
		imagen := model.Imagen{

			VarianteProducto: varianteId,
			Path:             *img,
			Fecha:            appUtils.FechaHoraBolivia(),
		}
		fmt.Println(imagen)
		s.ImagenRepository.CrearImgen(ctx, &imagen)
	}
	data := map[string]string{
		"productoVariante": varianteId.Hex(),
	}
	return data, nil
}
