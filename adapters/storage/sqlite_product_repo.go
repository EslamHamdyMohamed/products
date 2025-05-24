package storage

import (
	"gorm.io/gorm"
	"products/domain"
)

type SqliteProductRepository struct {
	db *gorm.DB
}

func NewSqliteProductRepository(db *gorm.DB) domain.ProductRepository {
	return SqliteProductRepository{
		db: db,
	}
}

func (r SqliteProductRepository) Create(product domain.Product) (domain.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r SqliteProductRepository) Delete(productId string) error {
	if err := r.db.Delete(&domain.Product{}, "id = ? ", productId).Error; err != nil {
		return err
	}
	return nil
}

func (r SqliteProductRepository) GetProduct(productId string) (domain.Product, error) {
	var product domain.Product
	if err := r.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r SqliteProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
