package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javierMorales9/rideshare-go/internal/config"
	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/http/handler"
	"github.com/javierMorales9/rideshare-go/internal/http/middleware"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)

	// 1) Conexión a Postgres
	db.MustConnect(cfg.DSN)

	// 3) Inyectar clave JWT en el middleware
	middleware.JwtKey = []byte(cfg.JWTSecret)

	// 4) Arrancar Gin
	r := gin.Default()
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)

	api := r.Group("/api", middleware.AuthRequired)
	{
		api.GET("/me", func(c *gin.Context) {
			uid := c.GetUint("currentUserID")
			c.JSON(http.StatusOK, gin.H{"user_id": uid})
		})
	}

	log.Printf("🚀  API escuchando en :%s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
