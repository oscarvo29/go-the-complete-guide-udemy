package main

import (
	"fmt"

	"api-project/db"
	"api-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080") // localhost:8080
	if err != nil {
		fmt.Println(err)
		return
	}
}
