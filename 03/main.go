package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

var products = []product{
	{ID: "1", Name: "kniha", Price: 200, Amount: 2},
	{ID: "2", Name: "rohlik", Price: 3, Amount: 5},
	{ID: "3", Name: "vejce", Price: 10, Amount: 6},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// bere na vstupu .json file -> v repozitáři předem vytvořen .json file "body.json"
// funkční curl: curl localhost:8080/products --include --header "Content-Type: application/json" -d @body.json --request "POST"
func createProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductById(id string) (*product, error) {
	for i, b := range products {
		if b.ID == id {
			return &products[i], nil
		}
	}

	return nil, errors.New("Product not found")
}

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func buyProduct(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found."})
		return
	}

	if product.Amount <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product not available."})
		return
	}

	product.Amount -= 1
	c.IndentedJSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusOK, product)
		}
	}
}

func main() {
	router := gin.Default()
	fmt.Print("Starting the APP")

	router.GET("/products", getProducts)
	router.POST("/products", createProduct)
	router.GET("/products/:id", productById)
	router.PATCH("/buy", buyProduct)
	router.DELETE("/products/delete", deleteProduct)
	router.Run("localhost:8080")
}
