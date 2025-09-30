package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

var productList []Product

func getProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlightReq(w, r)
	sendData(w,productList,200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlightReq(w, r)
	var newProduct Product
	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newProduct) 
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w,newProduct,201)
}

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Welcome to kholabazar")
}
func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreFlightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /",http.HandlerFunc(root))
	mux.Handle("GET /products", http.HandlerFunc(getProduct))          // get all products
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct)) // create new product
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe("127.0.0.1:8080", mux)

}

func init() {
	p1 := Product{
		ID:          1,
		Name:        "Wireless Headphones",
		Image:       "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=300&h=200&fit=crop",
		Price:       99.99,
		Description: "High-quality wireless headphones with noise cancellation",
		Category:    "Electronics",
	}
	p2 := Product{
		ID:          2,
		Name:        "Cotton T-Shir",
		Image:       "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=300&h=200&fit=crop",
		Price:       19.99,
		Description: "Comfortable 100% cotton t-shirt in various colors",
		Category:    "Clothing",
	}
	p3 := Product{
		ID:          3,
		Name:        "Coffee Mug",
		Image:       "https://images.unsplash.com/photo-1520031473529-2c06dea61853?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Price:       19.99,
		Description: "Ceramic coffee mug perfect for your morning brew",
		Category:    "Home",
	}
	p4 := Product{
		ID:          4,
		Name:        "Running Shoes",
		Image:       "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=300&h=200&fit=crop",
		Price:       79.99,
		Description: "Lightweight running shoes for all terrains",
		Category:    "Sports",
	}
	productList = append(productList, p1)
	productList = append(productList, p2)
	productList = append(productList, p3)
	productList = append(productList, p4)
}
