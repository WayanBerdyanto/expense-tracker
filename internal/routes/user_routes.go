package routes

import (
	"expense-tracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	r.POST("/login", authHandler.LoginUser)
	r.POST("/register", authHandler.RegisterUser)
}
