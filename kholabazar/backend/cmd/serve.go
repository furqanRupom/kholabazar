package cmd

import (
	"kholabazar/config"
	"kholabazar/repo"
	"kholabazar/rest"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	"kholabazar/rest/handlers/user"
	middleware "kholabazar/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	middleware := middleware.NewMiddlewares(conf)
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo, conf)
	reviewHandler := review.NewHandler()
	server := rest.NewServer(
		conf,
		userHandler,
		productHandler,
		reviewHandler,
	)
	server.Start()
}
