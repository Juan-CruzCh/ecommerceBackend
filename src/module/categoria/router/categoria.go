package router

import (
	"ecommerceBackend/src/module/categoria/controller"
	"net/http"
)

func NewCategoriaRouter(mux *http.ServeMux, controller *controller.Categoria) {

	mux.HandleFunc("GET /api/categoria", controller.ListarCategoria)
	mux.HandleFunc("POST /api/categoria", controller.CrearCategoria)
}
