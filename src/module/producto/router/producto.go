package router

import (
	"ecommerceBackend/src/module/producto/controller"
	"net/http"
)

func NewProductoRouter(mux *http.ServeMux, controller *controller.Producto) {
	mux.HandleFunc("POST /api/producto", controller.CrearProducto)
}
