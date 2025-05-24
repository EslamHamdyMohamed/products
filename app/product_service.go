package app

import (
	"products/domain"
)

type ProductUseCase interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	DeleteProduct(productId string) error
	GetProduct(productId string) (domain.Product, error)
	GetProducts() ([]domain.Product, error)
}
type ProductService struct {
	ProductRepository domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) ProductService {
	return ProductService{repo}
}

func (s ProductService) CreateProduct(product domain.Product) (domain.Product, error) {
	product, err := s.ProductRepository.Create(product)
	return product, err
}

func (s ProductService) DeleteProduct(productId string) error {
	err := s.ProductRepository.Delete(productId)
	return err
}

func (s ProductService) GetProduct(productId string) (domain.Product, error) {
	product, err := s.ProductRepository.GetProduct(productId)
	return product, err
}

func (s ProductService) GetProducts() ([]domain.Product, error) {
	products, err := s.ProductRepository.GetProducts()
	return products, err
}
