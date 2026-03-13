package controller

import (
	"context"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/usuario/dto"
	"ecommerceBackend/src/module/usuario/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Usuario struct {
	usuarioService *service.Usuario
	Validate       *validator.Validate
}

func NewUsuarioController(service *service.Usuario) Usuario {
	return Usuario{
		usuarioService: service,
	}
}
func (c *Usuario) CrearUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.UsuarioDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	err = validate.Struct(body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}

	resultado, err := c.usuarioService.CrearUsuario(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusCreated, resultado)

}
func (s *Usuario) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resultado, err := s.usuarioService.ListarUsuarios(ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)
}

func (s *Usuario) Eliminar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var id string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	resultado, err := s.usuarioService.Eliminar(ID, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}
func (c *Usuario) ActualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var id string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	validate := validator.New()
	var body dto.UsuarioDto

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	err = validate.Struct(body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	resultado, err := c.usuarioService.ActualizarUsuario(ID, &body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}
