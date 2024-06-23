package database

import (
	"database/sql"
	"fmt"

	"github.com/mateusjose98/imersao17/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db}
}

func (c *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *CategoryDB) GetCategoryById(id string) (*entity.Category, error) {
	row := c.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id)
	var category entity.Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &category, nil
}

func (c *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	fmt.Println("Creating category db " + category.Name + category.ID)
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return category.ID, nil
}
