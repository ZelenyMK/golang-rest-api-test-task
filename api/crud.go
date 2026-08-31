package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var products = make(map[int]Product) // will replace with DB later
var ID int = 0

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
		return
	}

	p.ID = ID
	products[ID] = p
	ID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

	slog.Info("Product created\n")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
		return
	}

	p := products[idNum]
	json.NewEncoder(w).Encode(p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idNum, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		slog.Error("bad thing happened", "error", err1)
		return
	}

	var newP Product

	err2 := json.NewDecoder(r.Body).Decode(&newP)
	if err2 != nil {
		slog.Error("bad thing happened", "error", err2)
		return
	}

	products[idNum] = newP
	slog.Info("Product was updated")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
		return
	}

	delete(products, idNum)

	slog.Info("Product with deleted")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
	slog.Info("Got products\n")
}
