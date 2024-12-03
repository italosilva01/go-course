package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func main() {
	BuildDb()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello everyone"))
	})
	router.Get("/products/{id}", GetProductByIDHandler)
	router.Get("/products", SearchProductsHandler)
	router.Post("/products",CreateProductHandler)
	// router.Put("/products/{ìd}",UpdateProductHandler)
	// router.Delete("/products/{id}", DeleteProductHandler)

	http.ListenAndServe(":8081",router)
}

//endpoints

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request){
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	product, ok:= memoryDb[id]
	if !ok {
		err := errors.New("product_not_found")
		http.Error(w,err.Error(),http.StatusNotFound)
		return
	}

	WriteJsonResponse(w,product, http.StatusOK)
}

//o http.ResponseWriter é responsável por construir e enviar 
// resposta http para o cliente
func SearchProductsHandler(w http.ResponseWriter, r *http.Request){
	productType := DecodeTypeQueryString(r);
	var matchedValues []*Product
	for _, value := range memoryDb {
		if value.Type == productType{
			matchedValues = append(matchedValues, value);
		}
	}
	WriteJsonResponse(w, matchedValues, http.StatusOK)
}
func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type");
}

func CreateProductHandler( w http.ResponseWriter, r *http.Request){
	product, err := DecodeProductFromBody(r);

	if err !=nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	idString := id.String()
	product.ID = idString
	memoryDb[idString] = product

	WriteJsonResponse(w,product, http.StatusCreated)
}
func DecodeProductFromBody(r *http.Request) (*Product, error){
	createProduct := &Product{}
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		return nil, err
	}

	return createProduct, nil
}


func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

func WriteJsonResponse(w http.ResponseWriter, obj interface{}, status int) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "JSON")
	w.WriteHeader(status)
	w.Write(bytes)
}