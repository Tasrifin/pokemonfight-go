package config

import (
	"fmt"
	"log"

	"github.com/Tasrifin/pokemonfight-go/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", constants.DB_USER, constants.DB_PASS, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")

	return db

}
