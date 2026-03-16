package service

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	productoTallaRepository "ecommerceBackend/src/module/producto/repository"
	"ecommerceBackend/src/module/stock/dto"
	"ecommerceBackend/src/module/stock/model"
	stockRepository "ecommerceBackend/src/module/stock/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Stock struct {
	stockRepository         stockRepository.Stock
	productoTallaRepository productoTallaRepository.ProductoTalla
}

func NewStockService(stockRepository stockRepository.Stock,
	productoTallaRepository productoTallaRepository.ProductoTalla) Stock {
	return Stock{
		stockRepository:         stockRepository,
		productoTallaRepository: productoTallaRepository,
	}
}
func (s *Stock) CrearStock(body *dto.DataStockDto, ctx context.Context) error {
	for _, v := range body.Stock {

		var productoTallaId bson.ObjectID

		tallaProducto, err := s.productoTallaRepository.BuscarProductoTalla(v.Producto, v.Talla, ctx)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				resultado, _ := s.productoTallaRepository.CrearProductoTalla(v.Producto, v.Talla, ctx)
				productoTallaId = resultado.InsertedID.(bson.ObjectID)
				continue
			}
			return err
		} else {
			productoTallaId = tallaProducto.ID
		}

		stock, err := s.stockRepository.VerificarStock(&v.Producto, &productoTallaId, ctx)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				data := model.Stock{
					Cantidad:      v.Cantidad,
					ProductoTalla: productoTallaId,
					Producto:      v.Producto,
					Fecha:         utils.FechaHoraBolivia(),
					Flag:          enum.FlagNuevo,
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
