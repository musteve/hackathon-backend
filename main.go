package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/hundler"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	http.HandleFunc("/hello", handlerHelloWorld)
	http.HandleFunc("/demo", demoHundler)
	http.HandleFunc("/book", bookHundler)
	http.ListenAndServe(":8080", nil)
}

func bookHundler(w http.ResponseWriter, r *http.Request) {
	hundler.BookHundler(w, r, db)
}

func demoHundler(w http.ResponseWriter, r *http.Request) {
	hundler.DemoHundler(w, r, db)
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
