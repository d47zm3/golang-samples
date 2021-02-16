package router

import (
	"cloudy/app/app"
	"cloudy/app/handler"

	"cloudy/app/router/middleware"

	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", handler.NewHandler(a.HandleReady, l))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)
		// Routes for books
		r.Method("GET", "/books", handler.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", handler.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", handler.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", handler.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", handler.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
