package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(router *gin.Engine) {
    router.GET("/transactions", controllers.GetTransactions)
    router.GET("/transactions/:id", controllers.GetTransactionByID)
    router.POST("/transactions", controllers.AddTransaction)
}