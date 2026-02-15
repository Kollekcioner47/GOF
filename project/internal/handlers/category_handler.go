package handlers

import (
	"net/http"

	"github.com/kollekcioner47/finance-app/internal/service"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(cs *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: cs}
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r) // нужно реализовать функцию получения userID из контекста
	categories, err := h.categoryService.GetUserCategories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Categories": categories,
	}
	renderTemplate(w, "categories/list.html", data)
}

func (h *CategoryHandler) CreateForm(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "categories/create.html", nil)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	name := r.FormValue("name")
	catType := r.FormValue("type") // "income" or "expense"
	_, err := h.categoryService.CreateCategory(userID, name, catType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

// вспомогательная функция для извлечения userID из контекста
func getUserID(r *http.Request) int {
	return r.Context().Value("userID").(int)
}
