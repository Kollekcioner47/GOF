package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kollekcioner47/finance-app/internal/service"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
	categoryService    *service.CategoryService
}

func NewTransactionHandler(ts *service.TransactionService, cs *service.CategoryService) *TransactionHandler {
	return &TransactionHandler{transactionService: ts, categoryService: cs}
}

func (h *TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	transactions, err := h.transactionService.GetUserTransactions(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Transactions": transactions,
	}
	renderTemplate(w, "transactions/list.html", data)
}

func (h *TransactionHandler) CreateForm(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	categories, err := h.categoryService.GetUserCategories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Categories": categories,
		"Today":      time.Now().Format("2006-01-02"),
	}
	renderTemplate(w, "transactions/create.html", data)
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	description := r.FormValue("description")

	_, err := h.transactionService.CreateTransaction(userID, categoryID, amount, description, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/transactions", http.StatusSeeOther)
}
