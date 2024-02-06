package handlers

import (
	"log"
	"microservices-in-go/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProducts(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}
