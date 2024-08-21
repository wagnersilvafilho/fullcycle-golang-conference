package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) GetByID(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	u, err := Get(h.db, int64(id))
	if err != nil {
		// TODO: validar se erro é por não existir nenhum registro
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Get(db *sql.DB, id int64) (*User, error) {
	stmt := `SELECT * FROM "users" WHERE id=$1`
	row := db.QueryRow(stmt, id)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Login,
		&u.Password,
		&u.CreatedAt,
		&u.ModifiedAt,
		&u.Deleted,
		&u.LastLogin,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
