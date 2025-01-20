package controllers

import (
	"net/http"
	"strconv"
	// "sync"
	"fmt"
	"myapp/database"
	"myapp/models"
	

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

// var mutex = &sync.Mutex{}

func CreateCustomer(c *gin.Context) {
	
	var customer models.Customer


	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Create a new context from the Gin request context
	ctx := c.Request.Context()

	collection := database.GetCollection("customers","customers")
	customer.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, customer) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
	fmt.Println(customer)
}







func GetCustomer(c *gin.Context) {
	ctx := c.Request.Context()

	// Use context with MongoDB operations
	collection := database.GetCollection("customers","customers")
	cursor, err := collection.Find(ctx, bson.M{}) // Pass context here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch customers"})
		return
	}
	defer cursor.Close(ctx)

	var customers []models.Customer
	for cursor.Next(ctx) {
		var customer models.Customer
		if err := cursor.Decode(&customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode customers"})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)

	
}



func UpdateCustomer(c *gin.Context) {
	var updatedItem models.Customer
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

}




func DeleteCustomer(c *gin.Context) {
	idParam := c.Query("id")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

}
