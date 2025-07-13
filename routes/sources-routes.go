package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSourceRoutes(router *gin.Engine) {
    router.GET("/sources", controllers.GetSources)
    router.GET("/sources/:id", controllers.GetSourceByID)
    router.POST("/sources", controllers.AddSources)
    router.PUT("/sources/:id", controllers.UpdateSource)
    router.DELETE("/sources/:id", controllers.DeleteSource)
}