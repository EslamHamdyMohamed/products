package domain

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type ProductRepository interface {
	Create(product Product) (Product, error)
	Delete(productId string) error
	GetProduct(productId string) (Product, error)
	GetProducts() ([]Product, error)
}
