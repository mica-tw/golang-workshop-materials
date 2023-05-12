package repository

import (
	"gorm.io/gorm"
	"tc-demo/internal/app/model"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetProduct(id uint) (*model.Product, error)
	UpdateProduct(id uint, product *model.Product) error
	DeleteProduct(id uint) error
	GetProducts() ([]*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProduct(id uint) (*model.Product, error) {
	product := &model.Product{}
	err := r.db.First(product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(id uint, product *model.Product) error {
	return r.db.Model(product).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *productRepository) GetProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
