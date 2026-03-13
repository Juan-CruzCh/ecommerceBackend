package controller

import (
	"ecommerceBackend/src/module/venta/service"
)

type Venta struct {
	ventaService *service.Venta
}

func NewVentaController(service *service.Venta) Venta {
	return Venta{
		ventaService: service,
	}
}
