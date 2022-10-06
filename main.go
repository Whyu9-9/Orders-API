package main

import (
	"order-api/databases"
	"order-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := databases.StartDB()
	g := handlers.New(db)

	router := gin.Default()

	router.GET("/orders", g.GetOrders)
	router.POST("/orders", g.CreateOrder)
	router.GET("/order/:id", g.GetOrder)
	router.PUT("/order/:id", g.UpdateOrder)
	router.DELETE("/order/:id", g.DeleteOrder)

	router.Run(":3000")
}
