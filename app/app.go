package app

import (
	"books-list/handlers"
	"books-list/logger"
	bookRepository "books-list/repository/book"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func StartApp() {
	db := ConnectDB()
	bookRepositoryDb := bookRepository.NewBookRepositoryDb(db)
	bookHandler := handlers.NewBooksService(bookRepositoryDb)

	router := mux.NewRouter()

	router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	router.HandleFunc("/books", bookHandler.AddBook).Methods("POST")
	router.HandleFunc("/books", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.RemoveBook).Methods("DELETE")

	port := os.Getenv("PORT_NAME")
	address := os.Getenv("ADDRESS_NAME")

	log.Printf("Server is running at %s: %s", address, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

var db *sqlx.DB

func logFatal(message string, err error) {
	if err != nil {
		logger.Error(message + err.Error())
	}
}

func ConnectDB() *sqlx.DB {
	gotenv.Load()
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal("Error while parsing Db address: ", err)

	db, err = sqlx.Open("postgres", pgUrl)
	if err != nil {
		logFatal("Error while opening DB: ", err)
	}

	err = db.Ping()
	if err != nil {
		logFatal("Error while pinging the database: ", err)
	}
	return db
}