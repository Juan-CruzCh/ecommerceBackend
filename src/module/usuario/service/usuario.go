package service

import "ecommerceBackend/src/module/usuario/repository"

type Usuario struct {
	usuarioRepository *repository.Usuario
}

func NewUsuarioService(repo *repository.Usuario) Usuario {
	return Usuario{
		usuarioRepository: repo,
	}
}
