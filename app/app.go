package app

import (
	"books-list/driver"
	"books-list/handlers"
	bookRepository "books-list/repository/book"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApp() {
	db := driver.ConnectDB()
	bookRepositoryDb := bookRepository.NewBookRepositoryDb(db)
	bookHandler := handlers.NewBooksService(bookRepositoryDb)

	router := mux.NewRouter()

	router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	router.HandleFunc("/books", bookHandler.AddBook).Methods("POST")
	router.HandleFunc("/books", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.RemoveBook).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}