package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type ProductoDto struct {
	Nombre      string        `json:"nombre" validate:"required"`
	Descripcion string        `json:"descripcion" validate:"required"`
	Categoria   bson.ObjectID `json:"categoria" validate:"required"`
	Destacado   *bool         `json:"destacado" validate:"required"`
	Publico     *bool         `json:"publico" validate:"required"`
}
