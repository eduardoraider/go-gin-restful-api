package database

import (
	"log"

	"github.com/eduardoraider/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	conn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&models.Student{})

}
