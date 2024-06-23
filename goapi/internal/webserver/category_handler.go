package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mateusjose98/imersao17/goapi/internal/entity"
	"github.com/mateusjose98/imersao17/goapi/internal/service"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(cs *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{cs}
}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	cat, error := wch.CategoryService.GetCategories()
	if error != nil {
		http.Error(w, "Error getting categories", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func (wch *WebCategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	if id == "" {
		http.Error(w, "Error getting category id", http.StatusBadRequest)
		return
	}
	cat, error := wch.CategoryService.GetCategoryById(id)
	if error != nil {
		http.Error(w, "Error getting category ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating category")
	var cat *entity.Category
	error := json.NewDecoder(r.Body).Decode(&cat)
	if error != nil {
		fmt.Println(error)
		http.Error(w, "Error decoding category", http.StatusBadRequest)
		return
	}
	id, error := wch.CategoryService.CreateCategory(cat)
	if error != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
