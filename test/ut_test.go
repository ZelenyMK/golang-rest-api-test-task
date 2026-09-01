package test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZelenyMK/golang-rest-api-test-task/api"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func TestAPI(t *testing.T) {
	slog.Info("Starting API tests")
	var router *chi.Mux = chi.NewRouter()
	router.Post("/products", api.CreateProduct)
	router.Get("/products/{id}", api.GetProduct)
	router.Put("/products/{id}", api.UpdateProduct)
	router.Delete("/products/{id}", api.DeleteProduct)
	router.Get("/products", api.GetProducts)

	slog.Info("Opening http test server")
	srv := httptest.NewServer(router)
	defer srv.Close()

	t.Run("CreateProduct", func(t *testing.T) {
		payload := `{"name": "Lego", "description": "brick game", "price": 10.10, "category": "toys"}`
		respons, err := http.Post(srv.URL+"/products", "application/json", strings.NewReader(payload))
		if err != nil {
			t.Fatalf("Error %v found in api.CreateProduct", err)
		}
		if respons.StatusCode != http.StatusCreated {
			t.Errorf("Wrong status code: %v", respons.StatusCode)
		}
	})
}
