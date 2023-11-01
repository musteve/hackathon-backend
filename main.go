package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid"
)


var db *sql.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, err")
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, err")
	}
	db = _db
}


type user struct {
	Id 	string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
	w.Header().Set("Access-Control-Allow-Origin", "*")


	rows, err := db.Query("select * from user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]user, 0)
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
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

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
	w.Header().Set("Access-Control-Allow-Origin", "*")


	length := r.ContentLength
	bytes := make([]byte, length)
	if _, err := r.Body.Read(bytes); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := new(user)
	if err := json.Unmarshal(bytes, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	newId, err := ulid.New(ms, entropy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(
		"insert into user (id, name, age) values (?, ?, ?)",
		newId.String(),
		body.Name,
		body.Age,
	)

	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	length := r.ContentLength
	bytes := make([]byte, length)
	if _, err := r.Body.Read(bytes); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := new(user)
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


type responseMessage struct {
	Message string `json:"message"`
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CITE_VERCEL"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bytes, err := json.Marshal(responseMessage{
		Message: "Hello, World!",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/hello", handlerHelloWorld)
	http.HandleFunc("/getusers", getUsers)
	http.HandleFunc("/adduser", addUser)
	http.HandleFunc("/deleteuser", deleteUser)
	http.ListenAndServe(":8080", nil)
}
