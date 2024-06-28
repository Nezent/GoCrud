package routes

import (
	"github.com/Nezent/GoCrud/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoute(router *gin.Engine) {
	router.GET("/", controllers.GetAccount)
	router.POST("/",controllers.CreateAccount)
	router.DELETE("/:id",controllers.DeleteAccount)
	router.PUT("/:id",controllers.UpdateAccount)
}