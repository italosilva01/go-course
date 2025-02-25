package productservices

import (
	"context"
	"first-tutorial/internal/product/productdomain/productentities"
	"first-tutorial/internal/product/productdomain/productrepositories"

	"github.com/google/uuid"
)

type productService struct {
	productRepositories *productrepositories.ProductRepository
}

func New() *productService {
	return &productService{
		productRepositories: productrepositories.New(),
	}
}

func (p *productService) Create(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	idString := id.String()
	product.ID = idString

	err = p.productRepositories.Create(ctx, product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) Search(ctx context.Context, productType string) ([]*productentities.Product, error) {
	return p.productRepositories.Search(ctx, productType)
}

func (p *productService) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	print("service id", id)
	return p.productRepositories.GetByID(ctx, id)
}

func (p *productService) Delete(ctx context.Context, id string) error {
	return p.productRepositories.Delete(ctx, id)
}

func (p *productService) Update(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	err := p.productRepositories.Update(ctx, product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
