package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mateusjose98/imersao17/goapi/internal/database"
	"github.com/mateusjose98/imersao17/goapi/internal/service"
	"github.com/mateusjose98/imersao17/goapi/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	productDB := database.NewProductDB(db)

	categoryService := service.NewCategoryService(*categoryDB)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/categories/{id}", webCategoryHandler.GetCategoryById)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Post("/categories", webCategoryHandler.CreateCategory)

	c.Get("/products/{id}", webProductHandler.GetProductById)
	c.Get("/products", webProductHandler.GetProducts)
	c.Post("/products", webProductHandler.CreateProduct)
	c.Get("/products/categories/{categoryID}", webProductHandler.GetProductsByCategory)
	fmt.Println("Server running on port 8080 >>>>>>>>>>>>>> ")
	http.ListenAndServe(":8080", c)

}
