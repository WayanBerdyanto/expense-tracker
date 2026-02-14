package main

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/handlers"
	"expense-tracker/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init DB
	db := config.ConnectDB()
	defer db.Close() // Tutup DB kalau aplikasi mati

	// 2. Init Handler (Inject DB ke Handler)
	expenseHandler := &handlers.ExpenseHandler{DB: db}
	authHandler := &handlers.AuthHandler{DB: db}

	// 3. Init Router (Gin)
	r := gin.Default()

	// 4. Routes
	v1 := r.Group("/api/v1")
	{
		routes.RegisterExpenseRoutes(v1, expenseHandler)
		routes.RegisterAuthRoutes(v1, authHandler)
	}

	// 5. Run Server
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
