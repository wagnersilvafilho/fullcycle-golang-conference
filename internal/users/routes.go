package users

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/auth"
)

var gh handler

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	gh = handler{db}

	r.Route("/users", func(r chi.Router) {

		r.Post("/", gh.Create)

		r.Group(func(r chi.Router) {
			r.Use(auth.Validate)

			r.Put("/{id}", gh.Modify)
			r.Delete("/{id}", gh.Delete)
			r.Get("/{id}", gh.GetByID)
			r.Get("/", gh.List)
		})
	})

}
