package database

import (
	"database/sql"
	"fmt"

	"github.com/mateusjose98/imersao17/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db}
}

func (p *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := p.db.Exec("INSERT INTO products (id, name, price, description, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Price, product.Descritpion, product.CategoryID, product.ImageURL)

	fmt.Println(product)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return product.ID, nil
}

func (p *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, price, description,category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Descritpion, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *ProductDB) GetProductById(id string) (*entity.Product, error) {
	row := p.db.QueryRow("SELECT id, name, price, description, category_id, image_url FROM products WHERE id = ?", id)
	var product entity.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Descritpion, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *ProductDB) GetProductsByCategory(categoryID string) ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, price, description,  category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Descritpion, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
