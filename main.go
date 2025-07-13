package main

import (
    "backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    routes.RegisterProductRoutes(router)
    routes.RegisterSourceRoutes(router)
	routes.RegisterTransactionRoutes(router)
    router.Run()
}