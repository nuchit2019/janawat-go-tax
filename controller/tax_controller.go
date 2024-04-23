package controller

import (
	// "github.com/nuchit2019/assessment-tax/config"
	"github.com/nuchit2019/assessment-tax/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TaxCalculations(c echo.Context) error {

	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:   "Tax Calculations",
	}

	return c.JSON(http.StatusOK, res)
}