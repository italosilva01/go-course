package producthttp

import (
	"errors"
	"first-tutorial/internal/encode"
	"first-tutorial/internal/product/productdb"
	"first-tutorial/internal/product/productdecode"
	"first-tutorial/internal/product/productdomain/productentities"
	"first-tutorial/internal/product/productdomain/productservices"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var productService = productservices.New()

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, ok := productdb.Memory[id]
	if !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	productType := productdecode.DecodeTypeQueryString(r)
	fmt.Println(productType)
	var matchedValues []*productentities.Product
	for _, value := range productdb.Memory {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}
	fmt.Println(matchedValues)

	encode.WriteJsonResponse(w, matchedValues, http.StatusOK)
}


func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := productdecode.DecodeProductFromBody(r)

	fmt.Println(product);
	fmt.Println("product");
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	idString := id.String()
	product.ID = idString
	productdb.Memory[idString] = product

	encode.WriteJsonResponse(w, product, http.StatusCreated)
}


