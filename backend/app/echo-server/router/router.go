package router

import (
	"mini-indobat/app/echo-server/handler"
	"net/http"

	"github.com/AMFarhan21/fres"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, productsHandler *handler.ProductsHandler, orderHandler *handler.OrderHandler) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, fres.Response.StatusOK("OK"))
	})

	products := e.Group("/products")
	products.GET("", productsHandler.ListOfProducts)

	products.POST("", productsHandler.AddProducts)

	order := e.Group("/order")
	order.POST("", orderHandler.OrderProducts)

}
