package rest

import (
	"fmt"
	"kholabazar/config"
	"kholabazar/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

func Start(conf config.Config) {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux, manager)
	addr := "127.0.0.1:" + strconv.Itoa(conf.HttpPort)
	fmt.Println("Server is running on port :", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
		os.Exit(1)
	}
}
