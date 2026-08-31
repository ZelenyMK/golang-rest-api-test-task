package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/ZelenyMK/golang-rest-api-test-task/DB"

	"github.com/go-chi/chi"
)

var ID int = 0

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
		return
	}
	p.ID = ID
	ID++
	p.Created_at = time.Now().UTC()
	DB.Database.Exec("INSERT INTO products (id, name, description, price, category, created_at) VALUES (?, ?, ?, ?, ?, ?)", p.ID, p.Name, p.Description, p.Price, p.Category, p.Created_at)
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

	var p Product = Product{}

	row := DB.Database.QueryRow("SELECT id, name, description, price, category, created_at FROM products WHERE id = ?", idNum)

	row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Created_at)
	json.NewEncoder(w).Encode(p)
	slog.Info("Got product\n")

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

	DB.Database.Exec("UPDATE products SET id = ?, name = ?, description = ?, price = ?, category = ?, created_at = ? WHERE id = ?",
		newP.ID, newP.Name, newP.Description, newP.Price, newP.Category, newP.Created_at, idNum)
	slog.Info("Product was updated")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("bad thing happened", "error", err)
		return
	}

	DB.Database.Exec("DELETE FROM products WHERE id = ?", idNum)

	slog.Info("Product was deleted")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var p Product = Product{}

	rows, _ := DB.Database.Query("SELECT id, name, description, price, category, created_at FROM products")
	for rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Created_at)
		json.NewEncoder(w).Encode(p)
	}
	slog.Info("Got products\n")
}
