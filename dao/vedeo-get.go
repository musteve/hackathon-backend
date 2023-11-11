package dao

import (
	"database/sql"
	"encoding/json"
	"main/model"
	"net/http"
)

func GetVedeos(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("select * from vedeo")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]model.Vedeo, 0)
	for rows.Next() {
		var u model.Vedeo
		if err := rows.Scan(
			&u.Id, 
			&u.Title, 
			&u.Author, 
			&u.Url,
			&u.Description,
			&u.Tag,
			&u.Last_update,
			&u.Last_update_date,
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}