package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
	psqlInfo string
)

func setCredentials() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")

	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

}

func connect() string {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(psqlInfo)
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(psqlInfo)
		panic(err)
	}

	return "Successfully Connected!"
}

func main() {
	setCredentials()
	fmt.Println(psqlInfo)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, World")
		fmt.Fprint(w, psqlInfo)
		fmt.Fprint(w, connect())
	})
	http.ListenAndServe(":8080", nil)
}
