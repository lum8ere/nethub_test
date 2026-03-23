package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func initRoutes() (*chi.Mux, error) {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With", "X-Request-Id", "X-Session-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// TODO: Добавить
	// 	r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("./swagger/doc.json"), // Use relative URL instead of absolute path
	// ))

	// TODO: add
	// r.Get("/healthz", auth.WithSimpleRestApiSmartContext(sctx, healthz.HealthzHandler))

	return r, nil
}
