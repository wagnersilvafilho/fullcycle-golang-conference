package folders

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/auth"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Route("/folders", func(r chi.Router) {
		r.Use(auth.Validate)

		r.Post("/", h.Create)
		r.Put("/", h.Modify)
		r.Get("/{id}", h.Get)
		r.Get("/", h.List)
		r.Delete("/{id}", h.Delete)
	})
}
