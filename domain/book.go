package domain

import (
	"books-list/dto"
	"books-list/err"
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
//type BookService interface {
//	GetBooks(w http.ResponseWriter, r *http.Request)
//	GetBook(w http.ResponseWriter, r *http.Request)
//	AddBook(w http.ResponseWriter, r *http.Request)
//	UpdateBook(w http.ResponseWriter, r *http.Request)
//	RemoveBook(w http.ResponseWriter, r *http.Request)
//}

// Secondary Port
//go:generate mockgen -destination=../mocks/repository/mockBookRepository.go -package=domain books-list/domain BookRepository
type BookRepository interface {
	GetBooks() ([]Book, *err.Error)
	GetBook(int) (*Book, *err.Error)
	AddBook(book Book) (int, error)
	UpdateBook(book Book) (int64, error)
	RemoveBook(id int) (int64, error)
}
