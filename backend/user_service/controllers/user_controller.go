package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserControllers struct {
	DB *sql.DB
}

type Claims struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Membership  string `json:"membership"`
	jwt.RegisteredClaims
}

// RegisterUser handles user registration
func (uc *UserControllers) RegisterUser(c *gin.Context) {
	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("[ERROR] Invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := "INSERT INTO users (id, name, email, password) VALUES (UUID(), ?, ?, ?)"
	_, err := uc.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("[ERROR] Failed to insert user into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	log.Println("[DEBUG] User registered successfully with Email:", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
func (uc *UserControllers) LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user struct {
		ID          string
		Name        string
		Email       string
		PhoneNumber string
		Membership  string
		Password    string
	}

	query := `SELECT id, name, email, phone_number, membership_tier, password 
	          FROM users WHERE email = ?`
	err := uc.DB.QueryRow(query, input.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Membership, &user.Password,
	)
	if err != nil || user.Password != input.Password {
		log.Println("[ERROR] Invalid login attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Println("[ERROR] JWT_SECRET is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	claims := Claims{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Membership:  user.Membership,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user_service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("[ERROR] Failed to sign token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

// GetUserDetails handles fetching user details
func (uc *UserControllers) GetUserDetails(c *gin.Context) {
	userID := c.Param("id")

	var user struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Membership  string `json:"membership"`
		CreatedAt   string `json:"created_at"`
	}

	query := "SELECT id, name, email, phone_number, membership_tier, created_at FROM users WHERE id = ?"
	err := uc.DB.QueryRow(query, userID).Scan(
		&user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Membership, &user.CreatedAt,
	)
	if err != nil {
		log.Println("[ERROR] User not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile allows users to update their profile
func (uc *UserControllers) UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")

	var input struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("[ERROR] Invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := "UPDATE users SET name = ?, email = ?, phone_number = ? WHERE id = ?"
	_, err := uc.DB.Exec(query, input.Name, input.Email, input.PhoneNumber, userID)
	if err != nil {
		log.Println("[ERROR] Failed to update user profile:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}
