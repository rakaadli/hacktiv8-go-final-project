package main

import (
	"fmt"
	"hacktiv8-final-project/config"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()
	route := gin.Default()

	fmt.Println(db)
	fmt.Println(route)

	route.Run(config.APP_PORT)

}
