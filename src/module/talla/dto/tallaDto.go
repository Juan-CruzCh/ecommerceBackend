package dto

type TallaDto struct {
	Nombre string `json:"nombre" validate:"required"`
}
