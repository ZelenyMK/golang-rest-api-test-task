package api

import (
	"encoding/json"
	"fmt"
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
		fmt.Printf("err: %v\n", err)
		return
	}

	p.ID = ID
	products[ID] = p
	ID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

	fmt.Printf("Product with id: %v created\n", p.ID)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	p := products[idNum]
	json.NewEncoder(w).Encode(p)

	fmt.Printf("%v\n", p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idNum, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		fmt.Printf("err: %v\n", err1)
		return
	}

	var newP Product

	err2 := json.NewDecoder(r.Body).Decode(&newP)
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
		return
	}

	products[idNum] = newP
	fmt.Printf("Product with id: %v was updated", idNum)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idNum, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		fmt.Printf("err: %v\n", err1)
		return
	}

	delete(products, idNum)

	fmt.Printf("Product with id: %v was deleted", idNum)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}
