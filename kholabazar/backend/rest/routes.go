package rest

import (
	"kholabazar/rest/handlers"
	"kholabazar/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {


	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthJWT,
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.AuthJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthJWT,
		),
	)
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)
	mux.Handle("POST /login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}
