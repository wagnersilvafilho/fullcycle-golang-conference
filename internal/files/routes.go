package files

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/auth"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/bucket"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/queue"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func SetRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Route("/files", func(r chi.Router) {
		r.Use(auth.Validate)

		r.Post("/", h.Create)
		r.Put("/", h.Modify)
		r.Delete("/{id}", h.Delete)
	})
}
