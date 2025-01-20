package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

func ProductRegisterRoutes(router *gin.Engine) {
	router.POST("/product/create", controllers.CreateProduct)
	router.GET("/product/get", controllers.GetProduct)
	router.PUT("/product/update", controllers.UpdateProduct)
	router.DELETE("/product/delete", controllers.DeleteProduct)
}
