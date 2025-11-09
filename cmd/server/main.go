package main

import (
	"URLShortener/internal/app"
	"URLShortener/internal/config"
	"URLShortener/internal/routers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env files found")
	}
	config := config.NewConfig()
	app := app.NewApp(config)

	r := routers.NewRouter(app)

	r.Run("localhost:8080")
}
