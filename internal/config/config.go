package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config agrupa todas las variables de entorno que necesitamos.
type Config struct {
	DSN       string // PostgreSQL DSN para Gorm y migrate
	Port      string // Puerto donde levantamos Gin
	JWTSecret string // Clave para firmar/validar JWT
}

// Load lee variables de entorno y aplica defaults sensatos.
func Load() Config {

	_ = godotenv.Load()

	cfg := Config{
		//DSN:       getenv("DSN", "host=localhost user=postgres dbname=rideshare sslmode=disable"),
		DSN:       getenv("DSN", ""),
		Port:      getenv("PORT", "8080"),
		JWTSecret: getenv("JWT_SECRET", "change-me"),
	}

	if cfg.JWTSecret == "change-me" {
		log.Println("[WARN] usando JWT_SECRET por defecto; cámbialo en producción")
	}
	return cfg
}

// helper privado
func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
