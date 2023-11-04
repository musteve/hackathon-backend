package dao

// import (
// 	"encoding/json"
// 	"main/model"
// 	"net/http"
// )

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("select * from user")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	users := make([]user.User, 0)
// 	for rows.Next() {
// 		var u user.User
// 		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		users = append(users, u)
// 	}

// 	bytes, err := json.Marshal(users)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(bytes)

// }