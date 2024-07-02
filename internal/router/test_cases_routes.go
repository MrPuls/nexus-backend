package router

import (
	"github.com/go-chi/chi/v5"
	"nexus/internal/handlers"
)

func SetupTestCaseRoutes(r chi.Router) {
	r.Route("/testcases", func(r chi.Router) {
		r.Post("/", handlers.CreateTestCase)
		r.Get("/", handlers.GetAllTestCases)
		r.Get("/search", handlers.SearchTestCases)
		r.Post("/bulk", handlers.TestCasesBulkOperation)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetTestCase)
			r.Put("/", handlers.UpdateTestCase)
			r.Delete("/", handlers.DeleteTestCase)
		})
	})
}
