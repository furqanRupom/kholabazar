package cmd

import (
	"fmt"
	"kholabazar/config"
	"kholabazar/infra/db"
	"kholabazar/repo"
	"kholabazar/rest"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	"kholabazar/rest/handlers/user"
	middleware "kholabazar/rest/middlewares"
	"os"
)

func Serve() {
	conf := config.GetConfig()
	dbCon, err := db.NewConnection(conf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if dbCon == nil {
		fmt.Println(err)
	}
	if dbCon != nil {
		fmt.Println("Database connected successfully!")
	}
	middleware := middleware.NewMiddlewares(conf)
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
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
