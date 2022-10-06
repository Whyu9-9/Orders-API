package handlers

import (
	"fmt"
	"net/http"
	"order-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func (h *Controller) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	h.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No order found!"})
	} else {
		result = gin.H{
			"status": "success",
			"data":   orders,
		}

		c.JSON(http.StatusOK, result)
	}

}

func (h *Controller) CreateOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.DB.Create(&order)
	result = gin.H{
		"message": "success create order",
		"data":    order,
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Controller) GetOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	id := ctx.Param("id")
	h.DB.Preload("Items").First(&order, id)
	if order.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No order found!"})
		return
	}

	result = gin.H{
		"status": "success",
		"data":   order,
	}

	ctx.JSON(http.StatusOK, result)
}
