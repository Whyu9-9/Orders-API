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

	router.Run(":3000")
}