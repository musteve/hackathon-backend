package hundler

// import (
// 	"main/dao"
// 	"net/http"
// )

// func DemoHundler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	switch r.Method {
// 	case http.MethodGet: dao.GetUsers(w, r)
// 	}
// }