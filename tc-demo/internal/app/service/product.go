package service

import (
	"tc-demo/internal/app/model"
	"tc-demo/internal/app/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProduct(id uint) (*model.Product, error)
	UpdateProduct(id uint, product *model.Product) error
	DeleteProduct(id uint) error
	GetProducts() ([]*model.Product, error)
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{repository: productRepository}
}

func (p *productService) CreateProduct(product *model.Product) error {
	return p.repository.CreateProduct(product)
}

func (p *productService) GetProduct(id uint) (*model.Product, error) {
	return p.repository.GetProduct(id)
}

func (p *productService) UpdateProduct(id uint, product *model.Product) error {
	return p.repository.UpdateProduct(id, product)
}

func (p *productService) DeleteProduct(id uint) error {
	return p.repository.DeleteProduct(id)
}

func (p *productService) GetProducts() ([]*model.Product, error) {
	return p.repository.GetProducts()
}
