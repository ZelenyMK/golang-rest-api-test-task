package main

import (
	"log/slog"
	"net/http"

	"github.com/ZelenyMK/golang-rest-api-test-task/api"
	"github.com/go-chi/chi"
)

func main() {
	var router *chi.Mux = chi.NewRouter()

	slog.Info("Starting API")

	router.Post("/products", api.CreateProduct)
	router.Get("/products/{id}", api.GetProduct)
	router.Put("/products/{id}", api.UpdateProduct)
	router.Delete("/products/{id}", api.DeleteProduct)
	router.Get("/products", api.GetProducts)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
	}

}
