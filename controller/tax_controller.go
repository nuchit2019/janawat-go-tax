package controller

import (
	// "github.com/nuchit2019/assessment-tax/config"
	"github.com/nuchit2019/assessment-tax/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TaxCalculations(c echo.Context) error {

	var req model.RequestModel

	if err := c.Bind(&req); err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to retrieve products...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	totalIncome := req.TotalIncome
	wht := req.WHT
	allowances := req.Allowances

	// Calculate Tax
	tax := totalIncome * wht
	for _, allallowances := range allowances {
		tax += allallowances.Amount
	}

	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    map[string]interface{}{"tax": tax},
	}

	return c.JSON(http.StatusOK, res)
}
