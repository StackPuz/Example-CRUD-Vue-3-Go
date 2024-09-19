package router

import (
	"app/controllers"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	productController := controllers.ProductController{}
	corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.Group("/api").
		GET("/products", productController.Index).
		POST("/products", productController.Create).
		GET("/products/:id", productController.Get).
		PUT("/products/:id", productController.Update).
		DELETE("/products/:id", productController.Delete)
	router.Run()
}
