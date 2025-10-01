package cmd

import (
	"fmt"
	"kholabazar/middleware"
	"kholabazar/router"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	InitRoutes(mux, manager)
	globalMux := router.GlobalRouter(mux)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe("127.0.0.1:8080", globalMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
	}
}
