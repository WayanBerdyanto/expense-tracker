package handlers

import (
	"database/sql"
	"expense-tracker/internal/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	DB *sql.DB
}

// 1. Create Expense
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var expense models.Expense

	// BindJSON mirip req.body di Express
	// Jika format JSON salah atau validasi gagal, return error
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query Raw SQL (Tanpa ORM)
	query := "INSERT INTO expenses (description, amount, category, created_at) VALUES (?, ?, ?, NOW())"
	result, err := h.DB.Exec(query, expense.Description, expense.Amount, expense.Category)

	// Explicit Error Handling (Go Way)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	expense.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success create expense",
		"data":    expense,
	})
}

// 2. Get All Expenses
func (h *ExpenseHandler) GetExpenses(c *gin.Context) {

	// Create Meta
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	// 1. Get total data
	var total int
	err := h.DB.QueryRow("SELECT COUNT(*) FROM expenses").Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	rows, err := h.DB.Query(`
		SELECT id, description, amount, category, created_at
		FROM expenses
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, perPage, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close() // Wajib tutup koneksi row setelah selesai (untuk memory leak prevention)

	var expenses []models.Expense

	// Looping baris database (seperti cursor)
	for rows.Next() {
		var e models.Expense
		if err := rows.Scan(
			&e.ID,
			&e.Description,
			&e.Amount,
			&e.Category,
			&e.CreatedAt,
		); err != nil {
			continue
		}
		expenses = append(expenses, e)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": expenses,
		"meta": gin.H{
			"total":       total,
			"page":        page,
			"per_page":    perPage,
			"total_pages": totalPages,
		},
	})
}
