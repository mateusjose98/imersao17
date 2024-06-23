package service

import (
	"fmt"

	"github.com/mateusjose98/imersao17/goapi/internal/database"
	"github.com/mateusjose98/imersao17/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(db database.CategoryDB) *CategoryService {
	return &CategoryService{db}
}

func (c *CategoryService) GetCategories() ([]*entity.Category, error) {
	return c.CategoryDB.GetCategories()
}

func (c *CategoryService) GetCategoryById(id string) (*entity.Category, error) {
	return c.CategoryDB.GetCategoryById(id)
}

func (c *CategoryService) CreateCategory(category *entity.Category) (string, error) {
	fmt.Println("Creating category service")
	cat := entity.NewCategory(category.Name)
	return c.CategoryDB.CreateCategory(cat)
}
