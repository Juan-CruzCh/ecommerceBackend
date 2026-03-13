package service

import "ecommerceBackend/src/module/cliente/repository"

type Cliente struct {
	clienteRepository *repository.Cliente
}

func NewClienteService(repo *repository.Cliente) Cliente {
	return Cliente{
		clienteRepository: repo,
	}
}
