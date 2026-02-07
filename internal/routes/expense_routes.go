package routes

import (
	"expense-tracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseRoutes(r *gin.RouterGroup, h *handlers.ExpenseHandler) {
	expense := r.Group("/expenses")
	{
		expense.POST("/", h.CreateExpense)
		expense.GET("/", h.GetExpenses)
		expense.GET("/:id", h.GetExpensesById)
		expense.PUT("/:id", h.UpdateExpense)
	}
}
