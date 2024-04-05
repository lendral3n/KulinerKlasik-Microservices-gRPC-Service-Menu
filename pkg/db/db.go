package db

import (
	"fmt"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(DB_USERNAME, DB_PASSWORD, DB_HOSTNAME string, DB_PORT int, DB_NAME string) Handler {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOSTNAME, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Menu{}, &models.StockDecreaseLog{})
	return Handler{db}
}
