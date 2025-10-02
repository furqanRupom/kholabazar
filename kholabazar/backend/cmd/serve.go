package cmd

import (
	"fmt"
	"kholabazar/middleware"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.PreFlight,middleware.Cors,middleware.Logger)

	fmt.Println("Server is running on port 8080")
	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux, manager)
	err := http.ListenAndServe("127.0.0.1:8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
	}
}
