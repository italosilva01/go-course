package producthttp

import (
	"first-tutorial/internal/encode"
	"first-tutorial/internal/product/productdecode"
	"first-tutorial/internal/product/productdomain/productservices"
	"fmt"
	"net/http"
)

var productService = productservices.New()

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	products, err := productService.Search(ctx, productType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encode.WriteJsonResponse(w, products, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productToCreate, err := productdecode.DecodeProductFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.Create(ctx, productToCreate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusCreated)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	product, err := productService.GetByID(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = productService.Delete(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, nil, http.StatusNoContent)

}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("r: ", r)
	id, err := productdecode.DecodeStringIDFromURI(r)
	fmt.Println("id: ", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	productToUpdate, err := productdecode.DecodeProductFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate.ID = id

	product, err := productService.Update(ctx, productToUpdate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)

}
