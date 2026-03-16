package controller

import (
	"context"
	"ecommerceBackend/src/core/utils"
	appUtils "ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/dto"
	"ecommerceBackend/src/module/producto/service"
	productoUtils "ecommerceBackend/src/module/producto/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Producto struct {
	productoService *service.Producto
	Validate        *validator.Validate
}

func NewProductoController(service *service.Producto, Validate *validator.Validate) Producto {
	return Producto{
		productoService: service,
		Validate:        Validate,
	}
}

func (c *Producto) CrearProducto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.ProductoDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.Validate.Struct(body)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := c.productoService.CrearProducto(&body, ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusCreated, resultado)

}

func (c *Producto) SubirImgenesProducto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": "productoVariante es obligatorio"})
		return
	}

	producto := r.FormValue("producto")
	productoId, err := appUtils.ValidadIdMongo(producto)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": "producto es obligatorio"})
		return
	}

	files := r.MultipartForm.File["imagenes"]
	if len(files) == 0 {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": "Debes subir al menos una imagen"})
		return
	}
	err = productoUtils.ValidarExtensiones(files)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.productoService.SubirImagenesProducto(productoId, files, ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "registrado"})
}

func (c *Producto) ListarProducto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	data, err := c.productoService.ListarProductos(ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusOK, data)
}

func (c *Producto) ListarImagenes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	varianteProducto := r.PathValue("varianteProducto")
	varianteProductoId, err := appUtils.ValidadIdMongo(varianteProducto)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	data, err := c.productoService.ListarImagenes(varianteProductoId, ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusOK, data)
}
func (c *Producto) ListarProductosPublico(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	destacado := r.URL.Query().Get("destacado")
	categoria := r.URL.Query().Get("categoria")

	data, err := c.productoService.ListarProductosPublico(destacado, categoria, ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusOK, data)
}

func (c *Producto) ListarProductosPublicoDetalle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	producto := r.PathValue("producto")
	productoId, err := utils.ValidadIdMongo(producto)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	data, err := c.productoService.ListarProductosPublicoDetalle(productoId, ctx)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	appUtils.ResponseJSON(w, http.StatusOK, data)
}
