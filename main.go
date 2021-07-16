package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"

	"database/sql"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book
var db *sql.DB

// инициализация переменных окружения
func init() {
	gotenv.Load()
}

// обработка ошибок
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	var book Book

	params := mux.Vars(r)
	rows := db.QueryRow("SELECT * FROM books WHERE id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)

}

func addBook(w http.ResponseWriter, r *http.Request) {

	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	db.QueryRow("INSERT INTO books (title, author, year) VALUES ($1, $2, $3);",
		book.Title, book.Author, book.Year)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	db.Exec("UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4", &book.Title, &book.Author, &book.Year, &book.ID)
	log.Println(book)
	json.NewEncoder(w).Encode(book)
}

func removeBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	params := mux.Vars(r)
	db.Exec("DELETE FROM books WHERE id = $1", params["id"])
}
