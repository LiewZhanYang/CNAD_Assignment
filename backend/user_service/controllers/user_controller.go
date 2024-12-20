package controllers

import (
	"database/sql"
	"fmt"
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
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// RegisterUser handles user registration
func (uc *UserControllers) RegisterUser(c *gin.Context) {
	var user struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO user (name, email, password, phoneNum, membership) VALUES (?, ?, ?, '', 'basic')"
	result, err := uc.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	userID, _ := result.LastInsertId()

	// Generate JWT token
	token, err := generateJWT(int(userID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": token})
}

// LoginUser handles user login and JWT creation
func (uc *UserControllers) LoginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user struct {
        ID          int
        Name        string
        Email       string
        PhoneNumber string
        Membership  string
        Password    string
    }

    query := "SELECT id, name, email, phoneNum, membership, password FROM user WHERE email = ?"
    err := uc.DB.QueryRow(query, input.Email).Scan(
        &user.ID, &user.Name, &user.Email, &user.PhoneNumber, &user.Membership, &user.Password,
    )
    if err != nil || user.Password != input.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":          user.ID,
        "name":        user.Name,
        "email":       user.Email,
        "phone_number": user.PhoneNumber,
        "membership":  user.Membership,
        "exp":         time.Now().Add(24 * time.Hour).Unix(), // Token expiry
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}


func (uc *UserControllers) UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")

	// Define the input structure
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		PhoneNum string `json:"phone_number" binding:"required"` // Map "phone_number" JSON field to "PhoneNum" struct field
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update query
	query := "UPDATE user SET name = ?, email = ?, phoneNum = ? WHERE id = ?"
	_, err := uc.DB.Exec(query, input.Name, input.Email, input.PhoneNum, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// GetAllUsers fetches all users
func (uc *UserControllers) GetAllUsers(c *gin.Context) {
	rows, err := uc.DB.Query("SELECT id, name, email, phoneNum, membership FROM user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var user struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Email      string `json:"email"`
			PhoneNum   string `json:"phoneNum"`
			Membership string `json:"membership"`
		}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNum, &user.Membership); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
			return
		}
		users = append(users, map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"phoneNum":   user.PhoneNum,
			"membership": user.Membership,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserDetails fetches details of a specific user by ID
func (uc *UserControllers) GetUserDetails(c *gin.Context) {
	userID := c.Param("id")
	var user struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		PhoneNum   string `json:"phoneNum"`
		Membership string `json:"membership"`
	}

	query := "SELECT id, name, email, phoneNum, membership FROM user WHERE id = ?"
	err := uc.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNum, &user.Membership)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Helper function to generate JWT token
func generateJWT(userID int, email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	claims := Claims{
		ID:    userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token valid for 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "CNAD_Assignment",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
