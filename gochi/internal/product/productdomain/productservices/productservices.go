package productservices

import (
	"context"
	"errors"
	"first-tutorial/internal/product/productdb"
	"first-tutorial/internal/product/productdomain/productentities"
)

type productService struct {

}

func New() *productService {
	return &productService{}
}

func (p *productService) GetByID (_ context.Context, id string) (*productentities.Product, error){
	product, ok := productdb.Memory[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}
	return product, nil
}