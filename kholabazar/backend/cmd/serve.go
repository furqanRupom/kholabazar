package cmd

import (
	"kholabazar/config"
	"kholabazar/rest"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	"kholabazar/rest/handlers/user"
)

func Serve() {
	conf := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()
	server := rest.NewServer(
		conf,
		userHandler,
		productHandler,
		reviewHandler,
	)
	server.Start()
}
