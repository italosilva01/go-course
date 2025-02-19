package producthttp

import (
	"first-tutorial/internal/encode"
	"first-tutorial/internal/product/productdecode"
	"first-tutorial/internal/product/productdomain/productservices"
	"net/http"
)

var productService = productservices.New()

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	println(r)
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
