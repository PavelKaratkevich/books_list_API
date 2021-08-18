package handlers

import (
	"books-list/domain"
	"books-list/dto"
	"books-list/err"
	"books-list/utils"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookHandlers struct {
	Repository domain.BookRepository
}

var book domain.Book
var error err.Error
var newBookRequest dto.NewBookRequest
var updateBookRequest dto.UpdateBookRequest

func (c *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []domain.Book
	var booksResponse []dto.BookResponse
	books, err := c.Repository.GetBooks(books)
	if err != nil {
		error.Message = "Server error"
		utils.SendError(w, http.StatusInternalServerError, error)
		return
	}
	for _, c := range books {
		booksResponse = append(booksResponse, c.ToDto())
	}
	utils.SendSuccess(w, booksResponse)
}

func (c BookHandlers) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var bookResponse dto.BookResponse
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
	bookResponse = book.ToDto()
	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(w, bookResponse)
}

func (c BookHandlers) AddBook(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&newBookRequest)
	if errs := newBookRequest.Validate(); errs != nil {
		utils.SendError(w, http.StatusBadRequest, *errs)
	} else {
		book := domain.Book{
			ID:     0,
			Title:  newBookRequest.Title,
			Author: newBookRequest.Author,
			Year:   newBookRequest.Year,
		}
		var bookResponse dto.NewBookResponse
		id, err := c.Repository.AddBook(book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		bookResponse.Id = id
		utils.SendSuccess(w, bookResponse)
	}
}

func (c BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&updateBookRequest)
	if updateBookRequest.Id == 0 || updateBookRequest.Author == "" || updateBookRequest.Title == "" || updateBookRequest.Year == "" {
		error.Message = "All fields should be filled in."
		utils.SendError(w, http.StatusBadRequest, error)
		return
	}
	book := domain.Book{
		ID:     updateBookRequest.Id,
		Title:  updateBookRequest.Title,
		Author: updateBookRequest.Author,
		Year:   updateBookRequest.Year,
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
	var error err.Error
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
