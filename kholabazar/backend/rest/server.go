package rest

import (
	"fmt"
	"kholabazar/config"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/user"
	"kholabazar/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	userHandler    *user.Handler 
	productHandler *product.Handler 
}

func NewServer(
	userHandler *user.Handler,
	productHandler *product.Handler,
) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start(conf config.Config) {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.ProductRoutes(mux, manager)
	addr := "127.0.0.1:" + strconv.Itoa(conf.HttpPort)
	fmt.Println("Server is running on port :", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
		os.Exit(1)
	}
}
