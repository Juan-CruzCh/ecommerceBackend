package service

import "ecommerceBackend/src/module/venta/repository"

type Venta struct {
	ventaRepository *repository.Venta
}

func NewVentaService(repo *repository.Venta) Venta {
	return Venta{
		ventaRepository: repo,
	}
}
