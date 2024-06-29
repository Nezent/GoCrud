package routes

import (
	"github.com/Nezent/GoCrud/controllers"
	"github.com/gin-gonic/gin"
)

var transfer_api_path = "/api/v1/transfer"

func TransferRoute(router *gin.Engine) {
	router.PUT(transfer_api_path + "/:id/from/:fromAccountID/to/:toAccountID", controllers.TransferBalanceController)
}