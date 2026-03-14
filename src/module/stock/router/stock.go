package router

import (
	"ecommerceBackend/src/module/stock/controller"
	"net/http"
)

func NewStockRouter(mux *http.ServeMux, controller *controller.Stock) {
	mux.HandleFunc("POST /api/stock", controller.CrearStock)
}
