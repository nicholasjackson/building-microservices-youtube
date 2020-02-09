package handlers

import (
	"net/http"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(&prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
