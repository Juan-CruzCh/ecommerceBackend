package src

import (
	"ecommerceBackend/src/core/config"
	categoriaRepository "ecommerceBackend/src/module/categoria/repository"
	clienteRepository "ecommerceBackend/src/module/cliente/repository"

	productoRepository "ecommerceBackend/src/module/producto/repository"
	stockRepository "ecommerceBackend/src/module/stock/repository"
	usuarioRepository "ecommerceBackend/src/module/usuario/repository"

	ventaRepository "ecommerceBackend/src/module/venta/repository"

	categoriaService "ecommerceBackend/src/module/categoria/service"
	clienteService "ecommerceBackend/src/module/cliente/service"
	productoService "ecommerceBackend/src/module/producto/service"
	stockService "ecommerceBackend/src/module/stock/service"
	usuarioService "ecommerceBackend/src/module/usuario/service"
	ventaService "ecommerceBackend/src/module/venta/service"

	categoriaController "ecommerceBackend/src/module/categoria/controller"
	clienteController "ecommerceBackend/src/module/cliente/controller"
	productoController "ecommerceBackend/src/module/producto/controller"
	stockController "ecommerceBackend/src/module/stock/controller"
	usuarioController "ecommerceBackend/src/module/usuario/controller"
	ventaController "ecommerceBackend/src/module/venta/controller"

	categoriaRouter "ecommerceBackend/src/module/categoria/router"
	clienteRouter "ecommerceBackend/src/module/cliente/router"
	productoRouter "ecommerceBackend/src/module/producto/router"
	stockRouter "ecommerceBackend/src/module/stock/router"
	usuarioRouter "ecommerceBackend/src/module/usuario/router"
	ventaRouter "ecommerceBackend/src/module/venta/router"

	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	categoriaRepository        categoriaRepository.Categoria
	clienteRepository          clienteRepository.Cliente
	productoRepository         productoRepository.Producto
	varianteProductoRepository productoRepository.VarianteProducto
	stockRepository            stockRepository.Stock
	usuarioRepository          usuarioRepository.Usuario
	ventaRepository            ventaRepository.Venta
	detalleVentaRepository     ventaRepository.DetalleVenta
	imagenRepository           productoRepository.Imagen
}
type App struct {
	ServerMux    *http.ServeMux
	Repositories *Repositories
	Validate     *validator.Validate
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		categoriaRepository:        categoriaRepository.NewCategoriaRepository(db),
		clienteRepository:          clienteRepository.NewClienteRepository(db),
		productoRepository:         productoRepository.NewProductoRepository(db),
		varianteProductoRepository: productoRepository.NewVarianteProductoRepository(db),
		stockRepository:            stockRepository.NewStockRepository(db),
		usuarioRepository:          usuarioRepository.NewUsuarioRepository(db),
		ventaRepository:            ventaRepository.NewVentaRepository(db),
		detalleVentaRepository:     ventaRepository.NewDetalleVentaRepository(db),
		imagenRepository:           productoRepository.NewImagenRepository(db),
	}
}

func NewApp(urlMongo string) *App {
	db, _, err := config.ConnectMongo(urlMongo, "kanna")
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	serverMux := http.NewServeMux()

	app := &App{
		ServerMux:    serverMux,
		Repositories: NewRepositories(db),
		Validate:     validate,
	}
	initCategoria(app)
	initProducto(app)
	initUsuario(app)
	return app
}

func (app *App) Run(port string) {
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	fmt.Println("Servidor corriendo en http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, app.ServerMux)
	if err != nil {
		log.Fatal(err)
	}
}

func initCategoria(app *App) {
	service := categoriaService.NewcategoriaService(app.Repositories.categoriaRepository)
	controller := categoriaController.NewCategoriaController(&service, app.Validate)
	categoriaRouter.NewCategoriaRouter(app.ServerMux, &controller)
}

func initCliente(app *App) {
	service := clienteService.NewClienteService(&app.Repositories.clienteRepository)
	controller := clienteController.NewClienteController(&service)
	clienteRouter.NewClienteRouter(app.ServerMux, &controller)
}

func initProducto(app *App) {
	service := productoService.NewProductoService(app.Repositories.productoRepository, app.Repositories.varianteProductoRepository, app.Repositories.imagenRepository)
	controller := productoController.NewProductoController(&service, app.Validate)
	productoRouter.NewProductoRouter(app.ServerMux, &controller)
}

func initStock(app *App) {
	service := stockService.NewStockService(&app.Repositories.stockRepository)
	controller := stockController.NewStockController(&service)
	stockRouter.NewStockRouter(app.ServerMux, &controller)
}

func initUsuario(app *App) {
	service := usuarioService.NewUsuarioService(app.Repositories.usuarioRepository)
	controller := usuarioController.NewUsuarioController(&service)
	usuarioRouter.NewUsuarioRouter(app.ServerMux, &controller)

}

func initVenta(app *App) {
	service := ventaService.NewVentaService(&app.Repositories.ventaRepository)
	controller := ventaController.NewVentaController(&service)
	ventaRouter.NewVentaRouter(app.ServerMux, &controller)

}
