package main

import (
	"net/http"
	"nethub-mdm/pkg/rest_middleware"
	"time"

	_ "nethub-mdm/docs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initRoutes(handlers *AppHandlers) (*chi.Mux, error) {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(rest_middleware.SessionID)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With", "X-Request-Id", "X-Session-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {

		r.Get("/swagger/*", httpSwagger.WrapHandler)

		r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})

		r.Route("/devices", func(r chi.Router) {
			r.Post("/", handlers.Device.Create)
			r.Get("/", handlers.Device.List)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.Device.GetByID)
				r.Put("/", handlers.Device.Update)
				r.Delete("/", handlers.Device.Delete)
			})
		})

		r.Route("/locations", func(r chi.Router) {
			r.Post("/", handlers.Location.Create)
			r.Get("/", handlers.Location.List)
			r.Route("/{id}", func(r chi.Router) {
				r.Put("/", handlers.Location.Update)
				r.Delete("/", handlers.Location.Delete)
			})
		})

		r.Route("/platforms", func(r chi.Router) {
			r.Get("/", handlers.Platform.List)
		})

		r.Get("/audit", handlers.Audit.List)
	})

	return r, nil
}
