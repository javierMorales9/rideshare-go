package main

import (
	"fmt"
	"log"

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
		api.GET("/me", handler.Me)

		// TripRequests
		api.POST("/trip_requests", handler.CreateTripRequest)
		api.GET("/trip_requests/:id", handler.ShowTripRequest)

		// Trips
		api.GET("/trips", handler.ListTrips)
		api.GET("/trips/:id", handler.ShowTrip)
		api.GET("/trips/:id/details", handler.TripDetails)
		api.GET("/trips/my", handler.ListMyTrips)
	}

	log.Printf("🚀  API escuchando en :%s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
