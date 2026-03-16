package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	appUtils "ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/dto"
	"ecommerceBackend/src/module/producto/model"
	"ecommerceBackend/src/module/producto/repository"
	productoUtils "ecommerceBackend/src/module/producto/utils"
	"errors"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Producto struct {
	productoRepository      repository.Producto
	productoTallaRepository repository.ProductoTalla
	imagenRepository        repository.Imagen
}

func NewProductoService(productoRepository repository.Producto, productoTallRepository repository.ProductoTalla, imagenRepository repository.Imagen) Producto {
	return Producto{
		productoRepository:      productoRepository,
		productoTallaRepository: productoTallRepository,
		imagenRepository:        imagenRepository,
	}
}

func (s *Producto) CrearProducto(producto *dto.ProductoDto, ctx context.Context) (map[string]string, error) {

	var body model.Producto = model.Producto{
		Nombre:       producto.Nombre,
		Descripcion:  producto.Descripcion,
		Categoria:    producto.Categoria,
		Destacado:    *producto.Destacado,
		Fecha:        appUtils.FechaHoraBolivia(),
		Publico:      *producto.Publico,
		Flag:         enum.FlagNuevo,
		PrecioCompra: producto.PrecioCompra,
		PrecioVenta:  producto.PrecioVenta,
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

func (s *Producto) SubirImagenesProducto(producto *bson.ObjectID, imagenes []*multipart.FileHeader, ctx context.Context) error {
	for _, v := range imagenes {

		img, err := productoUtils.GuardarImagen(v)
		if err != nil {

		}
		imagen := model.Imagen{
			Path:      *img,
			Fecha:     appUtils.FechaHoraBolivia(),
			Flag:      enum.FlagNuevo,
			Producto:  *producto,
			Principal: false,
		}

		s.imagenRepository.CrearImgen(ctx, &imagen)
	}

	return nil
}

func (s *Producto) ListarProductos(ctx context.Context) (*[]bson.M, error) {
	data, err := s.productoRepository.ListarProducto(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Producto) ListarImagenes(producto *bson.ObjectID, ctx context.Context) (*[]model.Imagen, error) {
	data, err := s.imagenRepository.ListarImagenes(producto, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Producto) ListarProductosPublico(descado string, categoria string, ctx context.Context) (*[]bson.M, error) {
	data, err := s.productoRepository.ListarProductosPublico(descado, categoria, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Producto) ListarProductosPublicoDetalle(producto *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	data, err := s.productoRepository.ListarProductosPublicoDetalle(producto, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
