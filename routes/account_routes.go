package routes

import (
	"github.com/Nezent/GoCrud/controllers"
	"github.com/gin-gonic/gin"
)

var resource_api_path = "/api/v1/resources"

func AccountRoute(router *gin.Engine) {
	router.GET(resource_api_path + "/", controllers.GetAccount)
	router.GET(resource_api_path + "/:id", controllers.GetAccountByID)
	router.POST(resource_api_path + "/",controllers.CreateAccount)
	router.DELETE(resource_api_path + "/:id",controllers.DeleteAccount)
	router.PUT(resource_api_path + "/:id",controllers.UpdateAccount)
}