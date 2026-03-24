package main

import (
	"net/http"
	"nethub-mdm/pkg/rest_middleware"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func initRoutes() (*chi.Mux, error) {

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

		r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})

		// r.Route("/devices", func(r chi.Router) {
		// 	r.Post("/", deviceHandler.Create)
		// 	r.Get("/", deviceHandler.List)

		// 	r.Route("/{id}", func(r chi.Router) {
		// 		r.Get("/", deviceHandler.GetByID)
		// 		r.Put("/", deviceHandler.Update)
		// 		r.Delete("/", deviceHandler.Delete)
		// 	})
		// })
	})
	// TODO: Добавить
	// 	r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("./swagger/doc.json"), // Use relative URL instead of absolute path
	// ))

	return r, nil
}
