package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

type Products struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := Products{
		Products: []Product{
			{
				ID:                 1,
				Title:              "iPhone 13 Pro Max",
				Description:        "An apple mobile which is nothing like apple",
				Price:              549,
				DiscountPercentage: 12.96,
				Rating:             4.69,
				Stock:              94,
				Brand:              "Apple",
				Category:           "smartphones",
				Thumbnail:          "https://cdn.dummyjson.com/product-images/1/thumbnail.jpg",
				Images: []string{
					"https://cdn.dummyjson.com/product-images/1/1.jpg",
					"https://cdn.dummyjson.com/product-images/1/2.jpg",
					"https://cdn.dummyjson.com/product-images/1/3.jpg",
					"https://cdn.dummyjson.com/product-images/1/4.jpg",
					"https://cdn.dummyjson.com/product-images/1/thumbnail.jpg",
				},
			},
			// Add more products here...
		},
		Total: 100,
		Skip:  0,
		Limit: 10,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8888", nil)
}