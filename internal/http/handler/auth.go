package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
	"github.com/javierMorales9/rideshare-go/internal/http/middleware"
	"github.com/javierMorales9/rideshare-go/internal/service"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// POST /auth/login
func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	user, err := service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString(middleware.JwtKey)

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"exp":      time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		"username": user.DisplayName(),
	})
}

// POST /auth/register
func Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name"  binding:"required"`
		Email     string `json:"email"      binding:"required,email"`
		Password  string `json:"password"   binding:"required"`
		Type      string `json:"type"       binding:"required,oneof=driver rider"`
		License   string `json:"drivers_license_number"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	u := &model.User{
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Email:                req.Email,
		Type:                 req.Type,
		DriversLicenseNumber: req.License,
	}
	if err := service.RegisterUser(u, req.Password); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": u.ID})
}

func Me(c *gin.Context) {
	uid := c.GetUint("currentUserID")
	c.JSON(http.StatusOK, gin.H{"user_id": uid})
}
