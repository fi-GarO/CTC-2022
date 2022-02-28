package main

import (
	"fmt"
	"log"
	"net/http"

)

type Product struct {
	Name string `json:"name"`
	Price float32 `json:"price"`
	Amount int16 `json:"amount"`
}

type Products []Product	

func getProduct() { // get product by name
}
func listProduct() {
	for index, item := range Products{} {
		println(index, item)
	}
} // list all products
func updateProduct() {}
func deleteProduct() {}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":4567", nil))
}

func main() {
    handleRequests()
}