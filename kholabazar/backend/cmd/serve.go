package cmd

import (
	"kholabazar/config"
	"kholabazar/rest"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/user"
)

func Serve() {
	conf := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server :=rest.NewServer(userHandler,productHandler)
	server.Start(conf)
}
