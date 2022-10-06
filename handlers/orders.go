package handlers

import (
	"net/http"
	"order-api/models"

	"github.com/gin-gonic/gin"
)

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&order)
	result = gin.H{
		"result": order,
	}

	ctx.JSON(http.StatusOK, result)
}
