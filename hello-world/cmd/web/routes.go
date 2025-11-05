package main

import (
	"net/http"

	"github.com/aishsanal/learning-go/hello-world/pkg/config"
	"github.com/aishsanal/learning-go/hello-world/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(printToConsole)
	mux.Use(NoSurf)
	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}
