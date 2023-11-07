package hundler

import (
	"database/sql"
	"main/dao"
	"net/http"
)

func DemoHundler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	switch r.Method {
	case http.MethodGet: dao.GetUsers(w, r, db)
	case http.MethodPost: dao.PostUser(w, r, db)
	case http.MethodDelete: dao.DeleteUser(w, r, db)
	case http.MethodPut: dao.PutUser(w, r, db)
	}
}