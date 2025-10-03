package product

import (
	"kholabazar/rest/middlewares"
	"net/http"
)

func (h *Handler) ProductRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthJWT,
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			middleware.AuthJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middleware.AuthJWT,
		),
	)

}
