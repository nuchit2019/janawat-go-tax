package controller

import (
	"github.com/nuchit2019/assessment-tax/config"
	"github.com/nuchit2019/assessment-tax/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	// Get the database instance
	db := configs.DB()
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to retrieve products...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}
	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    products,
	}
	return c.JSON(http.StatusOK, res)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	db := configs.DB()
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to retrieve product id:=" + id + "...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    product,
	}
	return c.JSON(http.StatusOK, res)
}
func CreateProduct(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		res := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to create product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	db := configs.DB()
	if err := db.Create(product).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to create product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := model.Response{
		Status:  http.StatusCreated,
		Message: "successfully",
		Data:    product,
	}
	return c.JSON(http.StatusCreated, res)

}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		res := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	existingProduct := model.Product{}
	if err := configs.DB().First(&existingProduct, id).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	if err := configs.DB().Model(&model.Product{}).Where("id =?", id).Updates(product).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	updatedProduct := model.Product{}
	if err := configs.DB().First(&updatedProduct, id).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    updatedProduct,
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	db := configs.DB()
	existingProduct := model.Product{}
	if err := configs.DB().First(&existingProduct, id).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to delete product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	if err := db.Delete(&model.Product{}, id).Error; err != nil {
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to delete product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := model.Response{
		Status:  http.StatusOK,
		Message: "Delete product successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, res)
}
