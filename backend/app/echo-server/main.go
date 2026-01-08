package main

import (
	"log"
	"mini-indobat/app/echo-server/handler"
	"mini-indobat/app/echo-server/router"
	"mini-indobat/repository"
	"mini-indobat/service/ordersService"
	"mini-indobat/service/productsService"
	"mini-indobat/utils/config"
	"mini-indobat/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.Load()

	db := database.GetDatabaseConnection(config.DBConnectionString)

	database.RunMigrations(config.DBConnectionString)

	// repository
	productsRepo := repository.NewProductsRepository(db)
	ordersRepo := repository.NewOrdersRepository(db)

	// service
	productsSvc := productsService.NewProductsService(productsRepo)
	ordersSvc := ordersService.NewOrdersService(ordersRepo, productsRepo)

	// handler
	productsHandler := handler.NewProductsHandler(productsSvc)
	orderHandler := handler.NewOrderHandler(ordersSvc)

	e := echo.New()
	e.Use(middleware.CORS())

	router.Router(e, productsHandler, orderHandler)

	log.Println("Successfully connected to the server")
	e.Logger.Fatal(e.Start(":" + config.Port))

}
