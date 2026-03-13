package controller

import (
	"ecommerceBackend/src/module/usuario/service"
)

type Usuario struct {
	usuarioService *service.Usuario
}

func NewUsuarioController(service *service.Usuario) Usuario {
	return Usuario{
		usuarioService: service,
	}
}
