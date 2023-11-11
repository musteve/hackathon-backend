package handler

import (
	"database/sql"
	"main/dao"
	"net/http"
)

func VedeoHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")

	switch r.Method {
	case http.MethodGet: dao.GetVedeos(w, r, db)
	case http.MethodPost: dao.InsertVedeo(w, r, db)
	case http.MethodDelete: dao.DeleteVedeo(w, r, db)
	case http.MethodPut: dao.PutVedeo(w, r, db)
	}
}