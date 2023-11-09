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

func InsertBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	length := r.ContentLength
	bytes := make([]byte, length)
	if _, err := r.Body.Read(bytes); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := new(model.Book)
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

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now()
	const layout = "2022-12-25"


	_, err = tx.Exec(
		"insert into book (id, title, author, issue_date, publisher, description, tag, last_update, last_update_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		newId.String(),
		body.Title,
		body.Author,
		body.Issue_date,
		body.Publisher,
		body.Description,
		body.Tag,
		newId.String(),
		t.Format(layout),
	)

	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}