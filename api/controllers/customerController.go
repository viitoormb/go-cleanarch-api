package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	api_helpers "github.com/viitoormb/go-cleanarch-api/api/helpers"
	"github.com/viitoormb/go-cleanarch-api/domain/customer"
	usecase "github.com/viitoormb/go-cleanarch-api/usecase/customer"
)

type CustomerController struct {
	useCase    customer.UseCase
	repository customer.Repository
}

func NewCustomerController(e *echo.Echo, repo customer.Repository) {
	handler := &CustomerController{
		useCase:    usecase.NewCustomerUseCase(repo),
		repository: repo,
	}

	e.GET("/customers", handler.getAll)
	e.POST("/customers/register", handler.registerNewCustomer)
}

func (h *CustomerController) registerNewCustomer(c echo.Context) error {
	request := &customer.CreateCustomer{}
	if err := c.Bind(request); err != nil {
		return err
	}

	output, err := h.useCase.CreateNewCustomer(c.Request().Context(), *request)

	if err != nil {
		res := api_helpers.HandleError(err, c.Request())
		return c.JSON(res.StatusCode, res)
	}

	return c.JSON(http.StatusOK, output)
}

func (h *CustomerController) getAll(c echo.Context) error {
	output, err := h.useCase.GetAllCustomers(c.Request().Context())

	if err != nil {
		res := api_helpers.HandleError(err, c.Request())
		return c.JSON(res.StatusCode, res)
	}

	return c.JSON(http.StatusOK, output)
}
