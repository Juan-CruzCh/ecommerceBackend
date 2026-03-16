package router

import (
	"ecommerceBackend/src/module/talla/controller"
	"net/http"
)

func NewTallaRouter(mux *http.ServeMux, controller *controller.Talla) {

	mux.HandleFunc("POST /api/talla", controller.CrearTalla)
	mux.HandleFunc("GET /api/talla", controller.ListarTalla)

}
