package main

import (
	"github.com/Nezent/GoCrud/config"
	"github.com/Nezent/GoCrud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.ConnectDB()
	routes.AccountRoute(router)
	router.Run(":8080")
}