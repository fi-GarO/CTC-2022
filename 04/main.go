package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Print("Database connected \n")
	}

	database.AutoMigrate(&product{})

	DB = database
}

type product struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

type CreateProductInput struct {
	Id     string `json:"id"`
	Name   string `json:"name" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type UpdateProductInput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product
	product := product{Id: input.Id, Name: input.Name, Price: input.Price, Amount: input.Amount}
	DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var newProduct product
	if err := DB.Where("id = ?", c.Param("id")).First(&newProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&newProduct).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": newProduct})
}

func FindProducts(c *gin.Context) {
	var products []product
	DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var newProduct product
	if err := DB.Where("id = ?", c.Param("id")).First(&newProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&newProduct)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func DeleteAllProducts(c *gin.Context) {

	var products []product
	DB.Delete(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
func main() {
	router := gin.Default()
	fmt.Print("Starting the APP\n")
	ConnectDatabase()
	router.GET("/products", FindProducts)
	router.POST("/products", CreateProduct)
	router.PATCH("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)
	router.DELETE("/products", DeleteAllProducts)

	router.Run(":8080")
}
