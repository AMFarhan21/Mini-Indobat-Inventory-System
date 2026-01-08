package handler

import (
	"log"
	"mini-indobat/models"
	"mini-indobat/service/productsService"
	"net/http"

	"github.com/AMFarhan21/fres"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	ProductsHandler struct {
		productsService productsService.Service
		validate        *validator.Validate
	}

	AddProductsInput struct {
		NamaObat string  `json:"nama_obat" validate:"required"`
		Stok     int     `json:"stok"`
		Harga    float64 `json:"harga"`
	}
)

func NewProductsHandler(productsService productsService.Service) *ProductsHandler {
	return &ProductsHandler{
		validate:        validator.New(),
		productsService: productsService,
	}
}

func (h *ProductsHandler) ListOfProducts(c echo.Context) error {
	ctx := c.Request().Context()

	products, err := h.productsService.GetAllProducts(ctx)
	if err != nil {
		log.Printf("Error on getting all products: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on getting all products"))
	}

	return c.JSON(http.StatusOK, fres.Response.StatusOK(products))
}

func (h *ProductsHandler) AddProducts(c echo.Context) error {
	ctx := c.Request().Context()
	var request AddProductsInput

	if err := c.Bind(&request); err != nil {
		log.Printf("Error on AddProducts request: %v", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	if err := h.validate.Struct(request); err != nil {
		log.Printf("Error on AddProducts request: %v", err.Error())
		return c.JSON(http.StatusBadRequest, fres.Response.StatusBadRequest(err.Error()))
	}

	product, err := h.productsService.CreateProduct(ctx, models.Products{
		NamaObat: request.NamaObat,
		Stok:     &request.Stok,
		Harga:    request.Harga,
	})
	if err != nil {
		log.Printf("Error on AddProducts internal: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, fres.Response.StatusInternalServerError("Error on Creating Products"))
	}

	return c.JSON(http.StatusCreated, fres.Response.StatusCreated(product))
}
