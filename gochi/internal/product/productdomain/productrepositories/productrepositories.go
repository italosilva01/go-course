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
	for _, value := range productdb.Memory {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *ProductRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	print("ID", id)
	product, ok := productdb.Memory[id]
	print("\n")
	print("product", product)
	if !ok {
		return nil, errors.New("product_not_found")
	}
	return product, nil
}

// Fazer Request de buscar produtos
