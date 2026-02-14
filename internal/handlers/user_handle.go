package handlers

import (
	"database/sql"
	"expense-tracker/internal/models"
	"expense-tracker/internal/res"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *sql.DB
}

var jwtSecret = []byte("awkawkoawkoawkoawkoa")

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	// Check if email already exists
	var existingID string
	err := h.DB.QueryRow("SELECT id FROM users WHERE email = ?", req.Email).Scan(&existingID)
	if err == nil {
		res.Error(c, http.StatusConflict, "EMAIL_EXISTS", "Email already registered")
		return
	} else if err != sql.ErrNoRows {
		res.Error(c, http.StatusInternalServerError, "DB_ERROR", err.Error())
		return
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "HASH_ERROR", "failed to hash password")
		return
	}

	userID := uuid.NewString()

	query := `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := h.DB.Exec(query, userID, req.Name, req.Email, string(hashed), time.Now(), time.Now())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "DB_ERROR", err.Error())
		return
	}
	if rowsAffected != 1 {
		res.Error(c, http.StatusInternalServerError, "DB_ERROR", "unexpected rows affected")
		return
	}

	if err != nil {
		res.Error(c, http.StatusInternalServerError, "DB_ERROR", err.Error())
		return
	}

	res.Success(c, http.StatusCreated, gin.H{
		"user_id":    userID,
		"name":       req.Name,
		"email":      req.Email,
		"created_at": time.Now(),
	}, nil)
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req models.User
	var hashedPassword string
	var userID string

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	err := h.DB.QueryRow(`
		SELECT id, password FROM users WHERE email = ?
	`, req.Email).Scan(&userID, &hashedPassword)

	if err != nil {
		res.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIAL", "email or password wrong")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		res.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIAL", "email or password wrong")
		return
	}

	token, _ := generateJWT(userID, req.Email)

	res.Success(c, http.StatusOK, gin.H{
		"user_id": userID,
		"token":   token,
	}, nil)
}

func generateJWT(userID string, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
