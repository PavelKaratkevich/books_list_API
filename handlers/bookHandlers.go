package handlers

import (
	"books-list/domain"
	"books-list/dto"
	"books-list/err"
	"books-list/utils"
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
var serverMessage utils.ServerMessage

func (c *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	var booksResponse []dto.BookResponse
	books, err := c.Repository.GetBooks()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, *err)
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
	book, err := c.Repository.GetBook(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, *err)
		return
	}
	bookResponse = book.ToDto()
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
			//error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		bookResponse.Id = id
		utils.SendSuccess(w, bookResponse)
	}
}

func (c BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBookRequest dto.UpdateBookRequest
	json.NewDecoder(r.Body).Decode(&updateBookRequest)
	if errs := updateBookRequest.Validate(); errs != nil {
		utils.SendError(w, http.StatusBadRequest, *errs)
		return
	}
		book := domain.Book{
			ID:     updateBookRequest.Id,
			Title:  updateBookRequest.Title,
			Author: updateBookRequest.Author,
			Year:   updateBookRequest.Year,
		}
		rowsUpdated, _ := c.Repository.UpdateBook(book)
		if rowsUpdated == 1 {
			serverMessage.Message = "Updated successfully"
			utils.SendSuccess(w, serverMessage)
		} else {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
	}

func (c BookHandlers) RemoveBook(w http.ResponseWriter, r *http.Request) {
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
	} else if rowsDeleted == 1 {
		serverMessage.Message = "Deleted successfully"
		utils.SendSuccess(w, serverMessage)
	}
}

func NewBooksService(repository domain.BookRepository) BookHandlers {
	return BookHandlers{
		repository,
	}
}