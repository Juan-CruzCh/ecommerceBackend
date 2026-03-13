package dto

type CategoriaDto struct {
	Nombre string `json:"nombre" validate:"required"`
}
