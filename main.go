package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"myapp/routes"
	"myapp/workers"
	"myapp/database"
)

func main() {
	
	go workers.StartWorkerPool(4)
	database.Connect()

	router := gin.Default()

	
	routes.ProductRegisterRoutes(router)
	routes.CustomerRegisterRoutes(router)

	
	fmt.Println("Server running on port 8081")
	router.Run(":8081")

	
	close(workers.TaskQueue)
	workers.WaitForWorkers()
}
