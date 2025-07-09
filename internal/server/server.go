package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	port := ":8080"

	e := echo.New()

	setupRoutes(e)

	log.Printf("starting server on %s", port)
	if err := e.Start(port); err != http.ErrServerClosed {
		log.Fatalf("fatal error: %v", err)
	}
}

func setupRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handlers.NewHealthHandler(e)
}
