package repository

import (
	"context"
	"ecommerceBackend/src/core/enum"
	"ecommerceBackend/src/core/utils"
	"ecommerceBackend/src/module/producto/model"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Producto interface {
	CrearProducto(ctx context.Context, producto *model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(ctx context.Context)
	EditarProducto(ctx context.Context)
	ListarProductosPublico(destacado string, categoria string, ctx context.Context) (*[]bson.M, error)
	ListarProductosPublicoDetalle(producto *bson.ObjectID, ctx context.Context) (*[]bson.M, error)
	ListarProducto(ctx context.Context) (*[]bson.M, error)

	countDocumentsProducto(ctx context.Context) (int, error)
}

type producto struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProductoRepository(db *mongo.Database) Producto {
	return &producto{
		db:         db,
		collection: db.Collection("Producto"),
	}
}

func (r *producto) CrearProducto(ctx context.Context, producto *model.Producto) (*mongo.InsertOneResult, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"nombre": producto.Nombre, "flag": enum.FlagNuevo})

	if cantidad > 0 {

		return nil, errors.New("El producto ya existe")
	}
	countDocumentsProducto, err := r.countDocumentsProducto(ctx)
	if err != nil {
		return nil, err
	}

	producto.Codigo = "PROD-" + strconv.Itoa(countDocumentsProducto)
	resultado, err := r.collection.InsertOne(ctx, producto)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
func (r *producto) countDocumentsProducto(ctx context.Context) (int, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return int(cantidad + 1), nil

}

func (r *producto) EliminarProducto(ctx context.Context) {

}

func (r *producto) EditarProducto(ctx context.Context) {

}

func (r *producto) ListarProducto(ctx context.Context) (*[]bson.M, error) {

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "flag", Value: enum.FlagNuevo,
					},
				},
			},
		},
		utils.Lookup("Categoria", "categoria", "_id", "categoria"),

		bson.D{
			{
				Key: "$project", Value: bson.D{
					{
						Key: "_id", Value: 1,
					},
					{
						Key: "codigo", Value: 1,
					},
					{
						Key: "nombre", Value: 1,
					},
					{
						Key: "categoria", Value: utils.ArrayElemAt("$categoria.nombre", 0),
					},
					{
						Key: "publico", Value: 1,
					},
					{
						Key: "destacado", Value: 1,
					},
					{
						Key: "precioCompra", Value: 1,
					},
					{
						Key: "precioVenta", Value: 1,
					},
					{
						Key: "descripcion", Value: 1,
					},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}
	return &data, nil
}

func (r *producto) ListarProductosPublico(destacado string, categoria string, ctx context.Context) (*[]bson.M, error) {
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "flag", Value: enum.FlagNuevo,
					},
					{
						Key: "publico", Value: true,
					},
				},
			},
		},
	}

	if destacado == "destacado" {
		pipeline = append(pipeline,
			bson.D{
				{
					Key: "$match", Value: bson.D{
						{
							Key: "destacado", Value: true,
						},
					},
				},
			},
		)
	}

	if categoria != "" {
		categoriId, err := utils.ValidadIdMongo(categoria)
		if err != nil {
			return nil, err
		}

		pipeline = append(pipeline,
			bson.D{
				{
					Key: "$match", Value: bson.D{
						{
							Key: "categoria", Value: categoriId,
						},
					},
				},
			},
		)
	}

	pipeline = append(pipeline,
		utils.Lookup("VarianteProducto", "_id", "producto", "varianteProducto"),
		utils.Unwind("$varianteProducto", false),
		bson.D{{
			Key: "$match", Value: bson.D{
				{
					Key: "varianteProducto.flag", Value: enum.FlagNuevo,
				},
			},
		}},
		utils.Lookup("Imagen", "varianteProducto._id", "varianteProducto", "imagen"),
		utils.Unwind("$imagen", false),
		bson.D{{
			Key: "$match", Value: bson.D{
				{
					Key: "imagen.flag", Value: enum.FlagNuevo,
				},
			},
		}},
		bson.D{
			{
				Key: "$group", Value: bson.D{
					{
						Key: "_id", Value: "$_id",
					},
					{
						Key: "precioVenta", Value: bson.D{{Key: "$first", Value: "$precioVenta"}},
					},
					{
						Key: "nombre", Value: bson.D{{Key: "$first", Value: "$nombre"}},
					},
					{
						Key: "imagen", Value: bson.D{{Key: "$first", Value: "$imagen.path"}},
					},
				},
			},
		},

		bson.D{
			{
				Key: "$project", Value: bson.D{
					{
						Key: "nombre", Value: 1,
					},
					{
						Key: "precioVenta", Value: 1,
					},
					{
						Key: "imagen", Value: 1,
					},
				},
			},
		},
	)

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}
	return &data, nil
}

func (r *producto) ListarProductosPublicoDetalle(producto *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "flag", Value: enum.FlagNuevo,
					},
					{
						Key: "_id", Value: producto,
					},
				},
			},
		},
		utils.Lookup("VarianteProducto", "_id", "producto", "varianteProducto"),
		utils.Lookup("Imagen", "varianteProducto.0._id", "varianteProducto", "imagen"),
		utils.Lookup("Stock", "_id", "producto", "stock"),

		bson.D{
			{
				Key: "$project", Value: bson.D{
					{
						Key: "nombre", Value: 1,
					},
					{
						Key: "descripcion", Value: 1,
					},
					{
						Key: "precioVenta", Value: 1,
					},
					{
						Key: "varianteProducto", Value: 1,
					},
					{
						Key: "imagen", Value: 1,
					},
					{
						Key: "stock", Value: 1,
					},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}

	return &data, nil
}
