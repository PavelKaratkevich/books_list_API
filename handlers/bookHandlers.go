package handlers

import (
	"books-list/domain"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"

	//"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type BookHandlers struct {
	Repository domain.BookRepository
}

var books []domain.Book
var book domain.Book
var error domain.Error

func (c *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := c.Repository.GetBooks(books)
	if err != nil {
		error.Message = "Server error"
		utils.SendError(w, http.StatusInternalServerError, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(w, books)
}

func (c BookHandlers) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	book, err := c.Repository.GetBook(book, id)
	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "Book ID Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		} else {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(w, book)
}

func (c BookHandlers) AddBook(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&book)
	if book.Author == "" || book.Title == "" || book.Year == "" {
		error.Message = "Enter missing fields."
		utils.SendError(w, http.StatusBadRequest, error) //400
		return
	}
	var bookWithId domain.Book
	id, err := c.Repository.AddBook(book)
	if err != nil {
		error.Message = "Server error"
		utils.SendError(w, http.StatusInternalServerError, error) //500
		return
	}
	bookWithId.ID = id
	utils.SendSuccess(w, bookWithId)
}

func (c BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&book)
	if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
		error.Message = "All fields are required."
		utils.SendError(w, http.StatusBadRequest, error)
		return
	}
	rowsUpdated, err := c.Repository.UpdateBook(book)
	if err != nil {
		error.Message = "Server error"
		utils.SendError(w, http.StatusInternalServerError, error) //500
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	utils.SendSuccess(w, rowsUpdated)
}

func (c BookHandlers) RemoveBook(w http.ResponseWriter, r *http.Request) {
	var error domain.Error
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	rowsDeleted, err := c.Repository.RemoveBook(id)

	if err != nil {
		error.Message = "Server error."
		utils.SendError(w, http.StatusInternalServerError, error) //500
		return
	}

	if rowsDeleted == 0 {
		error.Message = "Not Found"
		utils.SendError(w, http.StatusNotFound, error) //404
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	utils.SendSuccess(w, rowsDeleted)
}

func NewBooksService(repository domain.BookRepository) BookHandlers {
	return BookHandlers{
		repository,
	}
}
