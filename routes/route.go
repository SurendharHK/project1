package routes

import (
	"project1/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(route *gin.Engine, controller controllers.TransactionController) {
	route.POST("/banking/transfer", controller.CreatePayment)

}
