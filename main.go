package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/nuchit2019/assessment-tax/config"
	"github.com/nuchit2019/assessment-tax/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()
	//Middlware for logging
	e.Use(middleware.Logger())
	// Default route handler
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	// Middleware for Basic Authentication
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == configs.ApiAdmin() && password == configs.ApiAdminPass() {
			return true, nil
		}
		return false, nil
	}))

	// Connect to the database
	configs.InitDB()

	//********* test crud product
	// Group routes related to product endpoints
	// productRoute := e.Group("/product")
	// productRoute.GET("", controller.GetProducts)
	// productRoute.GET("/:id", controller.GetProduct)

	// productRoute.POST("", controller.CreateProduct)
	// productRoute.DELETE("/:id", controller.DeleteProduct)
	// productRoute.PUT("/:id", controller.UpdateProduct)

	tax := e.Group("/tax")
	tax.POST("/calculations", controller.TaxCalculations) // {HOST}/tax/calculations

	apiPort := configs.ApiPort()
	fmt.Printf("Server is running on port %s\n", apiPort)
	e.Logger.Fatal(e.Start(":" + apiPort))
}
