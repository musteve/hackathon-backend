package dao

import (
	"database/sql"
	"encoding/json"
	"main/model"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("select * from user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.LastUpdate); err != nil {
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