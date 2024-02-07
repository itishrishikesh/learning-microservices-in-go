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
		p.GetProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) PutProduct(writer http.ResponseWriter, request *http.Request) {
	// Todo: Add function here.
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//prod := data.Product{}
		//err := prod.FromJson(r.Body)
		//if err != nil {
		//	p.l.Println("[ERROR] deserializing product", err)
		//	http.Error(rw, "error reading product", http.StatusBadRequest)
		//}
		//ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		//r = r.WithContext(ctx)
		//next.ServeHTTP(rw, r)

		// Todo: Write function for this
	})
}
