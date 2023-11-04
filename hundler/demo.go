package hundler

import (
	"database/sql"
	"main/dao"
	"net/http"
)

func DemoHundler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet: dao.GetUsers(w, r, db)
	}
}