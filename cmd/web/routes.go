package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/KoVal177/site-02/pkg/config"
	"github.com/KoVal177/site-02/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
    mux := chi.NewRouter()

    mux.Use(middleware.Recoverer)
    mux.Use(NoSurf)
    mux.Use(SessionLoad)

    mux.Get("/", handlers.Repo.Home)
    mux.Get("/about", handlers.Repo.About)
    return mux
}
