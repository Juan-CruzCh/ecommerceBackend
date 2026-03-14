package controller

import (
	"context"
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

func (c *Producto) CrearVarianteProducto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.VarianteProductoDto
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

	resultado, err := c.productoService.CrearVarianteProducto(&body, ctx)
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
	producto := r.FormValue("producto")
	productoId, err := appUtils.ValidadIdMongo(producto)
	if err != nil {
		appUtils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": "Producto es obligatorio"})
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
	appUtils.ResponseJSON(w, http.StatusOK, map[string]string{"mensaje": "registrado"})
}
