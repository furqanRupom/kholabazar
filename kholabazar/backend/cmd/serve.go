package cmd

import (
	"fmt"
	"kholabazar/handlers"
	"kholabazar/router"
	"net/http"
)
func Serve(){
	mux := http.NewServeMux()
	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct))
	globalMux := router.GlobalRouter(mux)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe("127.0.0.1:8080", globalMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
	}
}