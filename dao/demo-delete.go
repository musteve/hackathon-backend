package dao

import (
	"database/sql"
	"encoding/json"
	"io"
	"main/model"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	_, err := db.Exec(
		"delete from user where id = ?",
		body.Id,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}