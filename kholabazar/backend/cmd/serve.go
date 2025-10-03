package cmd

import (
	"kholabazar/config"
	"kholabazar/rest"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	"kholabazar/rest/handlers/user"
	middleware "kholabazar/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	middleware := middleware.NewMiddlewares(conf)
	productHandler := product.NewHandler(middleware)
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
