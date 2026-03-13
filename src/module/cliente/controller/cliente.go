package controller

import (
	"ecommerceBackend/src/module/cliente/service"
)

type Cliente struct {
	clienteService *service.Cliente
}

func NewClienteController(service *service.Cliente) Cliente {
	return Cliente{
		clienteService: service,
	}
}