package hundler

import (
	"database/sql"
	"main/dao"
	"net/http"
)

func BookHundler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
	switch r.Method {
	case http.MethodGet: dao.GetBooks(w, r, db)
	case http.MethodPost: dao.InsertBook(w, r, db)
	case http.MethodDelete: dao.DeleteBook(w, r, db)
	case http.MethodPut: dao.PutBook(w, r, db)
	}
}