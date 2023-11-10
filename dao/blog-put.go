package dao

import (
	"database/sql"
	"encoding/json"
	"io"
	"main/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

func PutBlog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	length := r.ContentLength
	bytes := make([]byte, length)
	if _, err := r.Body.Read(bytes); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := new(model.Blog)
	if err := json.Unmarshal(bytes, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	newId, err := ulid.New(ms, entropy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now()

	_, err = db.Exec(
		"update blog set title = ? , author = ?, url = ?, description = ?, tag = ?, last_update = ?, last_update_date = ? where id = ?",
		body.Title,
		body.Author,
		body.Url,
		body.Description,
		body.Tag,
		newId.String(),
		t.String()[0:10],
		body.Id,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}