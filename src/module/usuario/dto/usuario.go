package dto

import "ecommerceBackend/src/core/enum"

type UsuarioDto struct {
	Ci        string `json:"ci"  validate:"required"`
	Nombre    string `json:"nombre"  validate:"required"`
	Celular   string `json:"celular"  validate:"required"`
	Apellidos string `json:"apellidos"  `

	Usuario  string  `json:"usuario"  validate:"required"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=8"`

	Rol enum.RolE `json:"rol"  validate:"required"`
}
