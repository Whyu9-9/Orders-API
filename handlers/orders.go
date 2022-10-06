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

	h.DB.Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)

}

func (h *Controller) CreateOrder(ctx *gin.Context) {
	var (
		order  models.Order
		item   models.Item
		result gin.H
	)

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&order)
	h.DB.Create(&item)
	result = gin.H{
		"result": order,
	}

	ctx.JSON(http.StatusOK, result)
}
