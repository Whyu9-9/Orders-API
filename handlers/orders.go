package handlers

import (
	"net/http"
	"order-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Controller) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	h.DB.Preload("Items").Find(&orders)

	if len(orders) <= 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No order found!",
		})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No order found!",
		})

		return
	}

	result = gin.H{
		"status": "success",
		"data":   order,
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Controller) UpdateOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	id := ctx.Param("id")
	h.DB.First(&order, id)

	if order.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No order found!",
		})

		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.DB.Save(&order)
	h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	result = gin.H{
		"message": "success update order",
		"data":    order,
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Controller) DeleteOrder(ctx *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	id := ctx.Param("id")
	h.DB.First(&order, id)

	if order.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No order found!",
		})

		return
	}

	h.DB.Delete(&order)
	//delete association
	h.DB.Model(&order).Association("Items").Clear()

	result = gin.H{
		"message": "success delete order",
	}

	ctx.JSON(http.StatusOK, result)
}
