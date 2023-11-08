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

func PutUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	length := r.ContentLength
	bytes := make([]byte, length)
	if _, err := r.Body.Read(bytes); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := new(model.User)
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

	_, err = db.Exec(
		"update user set name = ? , age = ?, last_update = ? where id = ?",
		body.Name,
		body.Age,
		newId.String(),
		body.Id,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}