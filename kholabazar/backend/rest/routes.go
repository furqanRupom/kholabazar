package rest

import (
	"kholabazar/rest/handlers"
	"kholabazar/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /",
		manager.With(
			http.HandlerFunc(handlers.Root),
		),
	)

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
}
