package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/stock/dto"
	"ecommerceBackend/src/module/stock/model"
	"ecommerceBackend/src/module/stock/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Stock struct {
	stockRepository repository.Stock
}

func NewStockService(repo repository.Stock) Stock {
	return Stock{
		stockRepository: repo,
	}
}
func (s *Stock) CrearStock(body *dto.DataStockDto, ctx context.Context) error {
	for _, v := range body.Stock {
		stock, err := s.stockRepository.VerificarStock(&v.Producto, &v.VarianteProducto, ctx)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				data := model.Stock{
					Cantidad:         v.Cantidad,
					VarianteProducto: v.VarianteProducto,
					Producto:         v.Producto,
					Fecha:            utils.FechaHoraBolivia(),
					Flag:             enum.FlagNuevo,
				}
				s.stockRepository.CrearStock(&data, ctx)
				continue
			}
			return err
		}
		var cantidad int = v.Cantidad + stock.Cantidad
		s.stockRepository.ActualizarStock(&stock.ID, cantidad, ctx)
	}

	return nil
}
