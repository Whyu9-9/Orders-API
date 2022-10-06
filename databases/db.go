package databases

import (
	"fmt"
	"log"
	"order-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "fga-db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "order-api"
)

func StartDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("cant open database ", err)
	}

	db.Debug().AutoMigrate(&models.Order{}, &models.Item{})

	return db
}
