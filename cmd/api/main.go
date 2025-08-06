package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/javierMorales9/rideshare-go/internal/config"
	"github.com/javierMorales9/rideshare-go/internal/db"
	//"github.com/javierMorales9/rideshare-go/internal/http/handler"
	"github.com/javierMorales9/rideshare-go/internal/http/middleware"
)

func main() {
	cfg := config.Load()

	// 1) Conexión a Postgres
	db.MustConnect(cfg.DSN)

	// 3) Inyectar clave JWT en el middleware
	middleware.JwtKey = []byte(cfg.JWTSecret)

	// 4) Arrancar Gin
	r := gin.Default()
	/* …rutas… */
	log.Fatal(r.Run(":" + cfg.Port))
}
