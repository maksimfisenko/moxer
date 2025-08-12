package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/maksimfisenko/moxer/internal/config"
	"github.com/maksimfisenko/moxer/internal/handlers"
	"github.com/maksimfisenko/moxer/internal/handlers/middleware"
	"github.com/maksimfisenko/moxer/internal/repo"
	"github.com/maksimfisenko/moxer/internal/repo/db"
	"github.com/maksimfisenko/moxer/internal/services"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func Start() {
	config.Load()

	e := echo.New()

	log.Print("connecting to db...")
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	setupRoutes(e, db)

	e.Use(middleware.RequestLogger())
	e.Use(middleware.JwtMiddleware())
	e.Use(echoMiddleware.CORS())

	log.Printf("starting server on %s...", config.Cfg.Port)
	if err := e.Start(config.Cfg.Port); err != http.ErrServerClosed {
		log.Fatalf("fatal error: %v", err)
	}
}

func setupRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	usersRepo := repo.NewUsersRepo(db)
	templatesRepo := repo.NewTemplatesRepo(db)

	authService := services.NewAuthSerice(usersRepo)
	templatesService := services.NewTemplatesService(templatesRepo)

	handlers.NewHealthHandler(e)
	handlers.NewAuthHandler(e, authService)
	handlers.NewTemplatesHandler(e, templatesService)
}
