package domain

import (
	"books-list/dto"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func (b Book) ToDto() dto.BookResponse {
	return dto.BookResponse{
		Id:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}

// Primary Port
type BookService interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	RemoveBook(w http.ResponseWriter, r *http.Request)
}

// Secondary Port
type BookRepository interface {
	GetBooks(books []Book) ([]Book, error)
	GetBook(book Book, id int) (Book, error)
	AddBook(book Book) (int, error)
	UpdateBook(book dto.UpdateBookRequest) (int64, error)
	RemoveBook(id int) (int64, error)
}
