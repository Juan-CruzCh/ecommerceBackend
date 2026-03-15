package router

import (
	"ecommerceBackend/src/module/producto/controller"
	"net/http"
)

func NewProductoRouter(mux *http.ServeMux, controller *controller.Producto) {
	fileServer := http.FileServer(http.Dir("./uploads"))
	mux.HandleFunc("POST /api/producto", controller.CrearProducto)
	mux.HandleFunc("GET /api/varianteProducto/{producto}", controller.ListarVarianteProducto)
	mux.HandleFunc("GET /api/producto", controller.ListarProducto)
	mux.HandleFunc("POST /api/varianteProducto", controller.CrearVarianteProducto)
	mux.HandleFunc("POST /api/imagenes", controller.SubirImgenesProducto)
	mux.HandleFunc("GET /api/imagenes/{varianteProducto}", controller.ListarImagenes)

	mux.Handle("GET /{filename}", fileServer)

}
