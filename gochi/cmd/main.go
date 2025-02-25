package main

import (
	"first-tutorial/internal/product/productdb"
	"first-tutorial/internal/product/producthttp"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	productdb.BuildDb()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello everyone"))
	})
	router.Get("/products/{id}", producthttp.GetProductByIDHandler)
	router.Get("/products", producthttp.SearchProductsHandler)
	router.Post("/products", producthttp.CreateProductHandler)
	router.Put("/products/{id}", producthttp.UpdateProductHandler)
	router.Delete("/products/{id}", producthttp.DeleteProductHandler)

	http.ListenAndServe(":8081", router)
}

//endpoints
