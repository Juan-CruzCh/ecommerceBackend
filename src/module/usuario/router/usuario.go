package router

import (
	"ecommerceBackend/src/module/usuario/controller"
	"net/http"
)

func NewUsuarioRouter(mux *http.ServeMux, controller *controller.Usuario) {
	mux.HandleFunc("POST /api/usuario", controller.CrearUsuarios)
	mux.HandleFunc("GET /api/usuario", controller.ListarUsuarios)
	mux.HandleFunc("DELETE /api/usuario/{id}", controller.Eliminar)
	mux.HandleFunc("PATCH /api/usuario/{id}", controller.ActualizarUsuarios)
}
