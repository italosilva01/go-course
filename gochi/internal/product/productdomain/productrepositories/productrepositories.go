package productrepositories

import (
	"context"
	"errors"
	"first-tutorial/internal/product/productdb"
	"first-tutorial/internal/product/productdomain/productentities"
)

type ProductRepository struct {
}

func New() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) Create(ctx context.Context, product *productentities.Product) error {
	productdb.Memory[product.ID] = product

	return nil
}

func (p *ProductRepository) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	var matchedValues []*productentities.Product
	print(productdb.Memory)
	for _, value := range productdb.Memory {
		print(value.Type)
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *ProductRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product, ok := productdb.Memory[id]

	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

// Fazer Request de buscar produtos
