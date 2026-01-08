package handler

import (
	"log"
	"mini-indobat/models"
	"mini-indobat/service/ordersService"
	"net/http"
	"strings"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	OrderHandler struct {
		validate      *validator.Validate
		ordersService ordersService.Service
	}

	OrderProductsInput struct {
		ProductId       int     `json:"product_id" validate:"required"`
		Quantity        int     `json:"quantity" validate:"required"`
		DiscountPercent float64 `json:"discount_percent"`
	}
)

func NewOrderHandler(ordersService ordersService.Service) *OrderHandler {
	return &OrderHandler{
		validate:      validator.New(),
		ordersService: ordersService,
	}
}

func (h *OrderHandler) OrderProducts(c echo.Context) error {
	var request OrderProductsInput

	if err := c.Bind(&request); err != nil {
		log.Printf("Error on OrderProducts request: %v", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on OrderProducts request: %v", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	order, err := h.ordersService.CreateOrder(models.Orders{
		ProductId:       request.ProductId,
		Quantity:        request.Quantity,
		DiscountPercent: request.DiscountPercent,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Printf("Error on OrderProducts, products isn't exists: %v", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		} else if strings.Contains(err.Error(), "stok tidak cukup") {
			log.Printf("Error on OrderProducts products stock: %v", err.Error())
			return c.JSON(http.StatusNotFound, fres.Response.StatusNotFound(err.Error()))
		}
		log.Printf("Error on OrderProducts internal: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on OrderProducts"))
	}

	return c.JSON(http.StatusCreated, fres.Response.StatusCreated(order))
}
