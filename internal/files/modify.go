package files

import (
	"database/sql"
	"time"
)

func Update(db *sql.DB, id int64, f *File) error {
	f.ModifiedAt = time.Now()
	stmt := `UPDATE "files" SET "name"=$1, "modified_at"=$2, "deleted"=$3 WHERE "id"=$4`
	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, f.Deleted, id)
	return err
}
