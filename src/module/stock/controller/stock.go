package controller

import (
	"context"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/stock/dto"
	"ecommerceBackend/src/module/stock/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	stockService *service.Stock
	Validate     *validator.Validate
}

func NewStockController(service *service.Stock, validate *validator.Validate) Stock {
	return Stock{
		stockService: service,
		Validate:     validate,
	}
}

func (c *Stock) CrearStock(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.DataStockDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.Validate.Struct(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.stockService.CrearStock(&body, ctx)
	fmt.Println(err)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "Registrado"})
}
