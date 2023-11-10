package handler

import (
	"database/sql"
	"main/dao"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
	switch r.Method {
	case http.MethodGet: dao.GetUsers(w, r, db)
	case http.MethodPost: dao.PostUser(w, r, db)
	case http.MethodDelete: dao.DeleteUser(w, r, db)
	case http.MethodPut: dao.PutUser(w, r, db)
	}
}