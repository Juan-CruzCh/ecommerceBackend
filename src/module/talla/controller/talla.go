package controller

import (
	"context"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/talla/dto"
	"ecommerceBackend/src/module/talla/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Talla struct {
	tallaService *service.Talla
	Validate     *validator.Validate
}

func NewTallaController(service *service.Talla, Validate *validator.Validate) Talla {
	return Talla{
		tallaService: service,
		Validate:     Validate,
	}
}

func (c *Talla) ListarTalla(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	resultado, err := c.tallaService.ListarTallas(ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}

func (c *Talla) CrearTalla(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.TallaDto

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

	resultado, err := c.tallaService.CrearTallas(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, resultado)

}
