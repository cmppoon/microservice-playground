package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"productsApi/data"
	"strconv"
)

func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	w.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, w)

	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get record by ID")
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}
	prod, err := data.GetProductByID(id)

	//if err != nil {
	//	p.l.Println("[ERROR] fetching product", err)
	//}
	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(prod, w)

	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
