package service

import (
	"fmt"

	"github.com/mateusjose98/imersao17/goapi/internal/database"
	"github.com/mateusjose98/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(db database.ProductDB) *ProductService {
	return &ProductService{db}
}

func (p *ProductService) GetProducts() ([]*entity.Product, error) {
	return p.ProductDB.GetProducts()
}

func (p *ProductService) GetProductById(id string) (*entity.Product, error) {
	return p.ProductDB.GetProductById(id)
}

func (p *ProductService) CreateProduct(name, description string, price float64, categoryID, imageURL string) (string, error) {
	fmt.Println("category >> ", categoryID)
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	return p.ProductDB.CreateProduct(product)
}

func (p *ProductService) GetProductsByCategory(categoryID string) ([]*entity.Product, error) {
	fmt.Println("categoryID >> ", categoryID)
	return p.ProductDB.GetProductsByCategory(categoryID)
}
