package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mateusjose98/imersao17/goapi/internal/entity"
	"github.com/mateusjose98/imersao17/goapi/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ps *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ps}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod, error := wph.ProductService.GetProducts()
	if error != nil {
		http.Error(w, "Error getting products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prod)
}

func (wph *WebProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Error getting product id", http.StatusBadRequest)
		return
	}
	prod, error := wph.ProductService.GetProductById(id)
	if error != nil {
		http.Error(w, "Error getting product ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prod)
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var prod *entity.Product
	error := json.NewDecoder(r.Body).Decode(&prod)
	if error != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}
	fmt.Println("Controller >> ", prod)
	id, error := wph.ProductService.CreateProduct(prod.Name, prod.Descritpion, prod.Price, prod.CategoryID, prod.ImageURL)
	if error != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}

func (wph *WebProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {

	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "Error getting category id", http.StatusBadRequest)
		return
	}
	prod, error := wph.ProductService.GetProductsByCategory(categoryID)
	if error != nil {
		http.Error(w, "Error getting products by category", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prod)
}
