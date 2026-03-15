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

func (s *Producto) CrearVarianteProducto(body *dto.VarianteProductoDto, ctx context.Context) (map[string]string, error) {
	variante := model.VarianteProducto{
		Talla:    body.Talla,
		Color:    body.Color,
		Producto: body.Producto,
		Fecha:    appUtils.FechaHoraBolivia(),
		Flag:     enum.FlagNuevo,
	}
	resultado, err := s.varianteProductoRepository.CrearVarianteProducto(ctx, &variante)
	if err != nil {
		return nil, err
	}
	varianteId, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("Error de parseo")
	}

	data := map[string]string{
		"productoVariante": varianteId.Hex(),
	}
	return data, nil
}
func (s *Producto) SubirImagenesProducto(producto *bson.ObjectID, variante *bson.ObjectID, imagenes []*multipart.FileHeader, ctx context.Context) error {
	for _, v := range imagenes {

		img, err := productoUtils.GuardarImagen(v)
		if err != nil {

		}
		imagen := model.Imagen{
			VarianteProducto: *variante,
			Path:             *img,
			Fecha:            appUtils.FechaHoraBolivia(),
			Flag:             enum.FlagNuevo,
			Producto:         *producto,
			Principal:        false,
		}

		s.ImagenRepository.CrearImgen(ctx, &imagen)
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

func (s *Producto) ListarVarianteProducto(producto *bson.ObjectID, ctx context.Context) (*[]model.VarianteProducto, error) {
	data, err := s.varianteProductoRepository.ListarVarianteProducto(producto, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Producto) ListarImagenes(variante *bson.ObjectID, ctx context.Context) (*[]model.Imagen, error) {
	data, err := s.ImagenRepository.ListarImagenes(variante, ctx)
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
