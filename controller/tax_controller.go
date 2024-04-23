package controller

import (
	// "github.com/nuchit2019/assessment-tax/config"
	"github.com/nuchit2019/assessment-tax/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TaxCalculations(c echo.Context) error {
    // Bind request data to struct
    req := model.RequestModel{}
    if err := c.Bind(&req); err != nil {
        return handleError(c, http.StatusNotFound, "Failed to retrieve products...Err:"+err.Error())
    }

    // Calculate tax-able income
    taxableIncome := calculateTaxableIncome(req.TotalIncome, req.Allowances)

    // Calculate tax
    tax := calculateTax(taxableIncome)

    // Prepare response
    res := model.Response{
        Status:  http.StatusOK,
        Message: "successfully",
        Data:    map[string]interface{}{"tax": tax},
    }

    return c.JSON(http.StatusOK, res)
}

func calculateTaxableIncome(totalIncome float64, allowances []model.Allowance) float64 {
    incomeAfterAllowance := totalIncome
    for _, allowance := range allowances {
        incomeAfterAllowance -= allowance.Amount
    }
    return incomeAfterAllowance
}

func calculateTax(taxableIncome float64) float64 {

    if taxableIncome <= 150000 {
        return 0
    } else if taxableIncome <= 500000 {
        return 29000
    } else if taxableIncome <= 1000000 {
        return 0
    } else {
        return 0
    }
	
}

func handleError(c echo.Context, statusCode int, message string) error {
    res := model.Response{
        Status:  statusCode,
        Message: message,
    }
    return c.JSON(statusCode, res)
}

