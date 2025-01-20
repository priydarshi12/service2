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

func CreateProduct(c *gin.Context) {
	
	var product models.Product


	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Create a new context from the Gin request context
	ctx := c.Request.Context()

	collection := database.GetCollection("products","products")
	product.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, product) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
	fmt.Println(product)
}







func GetProduct(c *gin.Context) {
	ctx := c.Request.Context()

	// Use context with MongoDB operations
	collection := database.GetCollection("products","products")
	cursor, err := collection.Find(ctx, bson.M{}) // Pass context here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode products"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)

	
}



func UpdateProduct(c *gin.Context) {
	var updatedItem models.Product
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

}




func DeleteProduct(c *gin.Context) {
	idParam := c.Query("id")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

}
