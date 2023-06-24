package main

import (
	"log"

	"github.com/Tasrifin/pokemonfight-go/config"
	"github.com/Tasrifin/pokemonfight-go/constants"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()

	log.Println(db)

	route.Run(constants.APP_PORT)
}
