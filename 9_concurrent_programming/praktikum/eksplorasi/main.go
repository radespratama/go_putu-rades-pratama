package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	products := make(chan Product)

	go getProducts(products)

	for product := range products {
		fmt.Printf("=========\nPRODUCT DATA\n=========\n")
		fmt.Printf("title: %s\n", product.Title)
		fmt.Printf("price: %.2f\n", product.Price)
		fmt.Printf("category: %s\n", product.Category)
		fmt.Printf("=========\n")
	}
}

type Product struct {
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Category string  `json:"category"`
}

func getProducts(products chan Product) {
	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var productData []Product
	err = json.Unmarshal(body, &productData)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, product := range productData {
		products <- product
	}
}
