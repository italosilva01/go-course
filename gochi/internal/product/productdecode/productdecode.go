package productdecode

import (
	"encoding/json"
	"errors"
	"first-tutorial/internal/product/productdomain/productentities"
	"net/http"

	"github.com/go-chi/chi"
)

func DecodeProductFromBody(r *http.Request) (*productentities.Product, error) {
	createProduct := &productentities.Product{}
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

func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}
