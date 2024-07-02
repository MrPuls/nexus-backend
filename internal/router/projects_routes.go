package router

import (
	"github.com/go-chi/chi/v5"
	"nexus/internal/handlers"
)

func SetupProjectRoutes(r chi.Router) {
	r.Route("/projects", func(r chi.Router) {
		r.Post("/", handlers.CreateProject)
		r.Get("/", handlers.GetAllProjects)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetProject)
			r.Put("/", handlers.UpdateProject)
			r.Delete("/", handlers.DeleteProject)
		})
	})
}
