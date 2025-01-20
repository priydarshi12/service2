package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

func CustomerRegisterRoutes(router *gin.Engine) {
	router.POST("/customer/create", controllers.CreateCustomer)
	router.GET("/customer/get", controllers.GetCustomer)
	router.PUT("/customer/update", controllers.UpdateCustomer)
	router.DELETE("/customer/delete", controllers.DeleteCustomer)
}
