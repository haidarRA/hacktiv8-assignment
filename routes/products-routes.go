package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func welcomeAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome!",
	})
}

func pingAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong!",
	})
}

func RegisterProductRoutes(router *gin.Engine) {
    router.GET("/", welcomeAPI)
    router.GET("/ping", pingAPI)
    router.GET("/products", controllers.GetProducts)
    router.GET("/products/:id", controllers.GetProductByID)
    router.POST("/products", controllers.AddProducts)
    router.PUT("/products/:id", controllers.UpdateProduct)
    router.DELETE("/products/:id", controllers.DeleteProduct)
}