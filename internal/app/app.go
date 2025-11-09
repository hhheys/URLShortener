package app

import (
	"URLShortener/internal/config"
	"URLShortener/internal/db"
	"URLShortener/internal/handler"
	"URLShortener/internal/repository"
	"URLShortener/internal/service"
	"database/sql"
)

type App struct {
	DB *sql.DB

	Repositories *Repositories
	Services     *Services
	Handlers     *Handlers
}

type Repositories struct {
	ShortenerRepository *repository.ShortenerRepository
}

type Services struct {
	ShortenerService *service.ShortenerService
}

type Handlers struct {
	ShortenerHandler *handler.ShortenerHandler
}

func NewApp(config *config.Config) *App {
	dbConnection := db.CreateConnection(config.DatabaseURL)

	repositories := NewRepositories(dbConnection)
	services := NewServices(repositories)
	handlers := NewHandlers(services)

	return &App{
		DB: dbConnection,

		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}
}

func NewRepositories(conn *sql.DB) *Repositories {
	return &Repositories{
		ShortenerRepository: repository.NewShortenerRepository(conn),
	}
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		ShortenerService: service.NewShortenerService(repositories.ShortenerRepository),
	}
}

func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		ShortenerHandler: handler.NewShortenerHandler(services.ShortenerService),
	}
}
