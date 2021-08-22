package bookRepository

import (
	domain2 "books-list/domain"
	//"books-list/handlers"
	domain "books-list/mocks/repository"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"testing"
)
var mockBookHandlers *domain.MockBookRepository
var router *mux.Router
var ctrl *gomock.Controller


func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockBookHandlers = domain.NewMockBookRepository(ctrl)
	router = mux.NewRouter()
}

func Test_AddBook_should_return_error_when_new_book_is_added (t *testing.T) {
	setup(t)
// Arrange
	book := domain2.Book{}
	mockBookHandlers.EXPECT().AddBook(book).Return(0, &appErr)

// Act
	_, err := mockBookHandlers.AddBook(book)

// Assert
	if err == nil {
		t.Error("Test failed while returning error")
	}
}

func Test_GetBooks_shoud_return_books (t *testing.T) {
	setup(t)
	books := []domain2.Book{{ID: 13,Title: "Pavel", Author: "book", Year: "2021"}}
	mockBookHandlers.EXPECT().GetBooks().Return(books, nil)

	_, err := mockBookHandlers.GetBooks()
	if err != nil {
		t.Error("Test failed when returning books")
	}
}

func Test_GetBooks_shoud_return_error (t *testing.T) {
	setup(t)

// Arrange
	mockBookHandlers.EXPECT().GetBooks().Return(nil, &appErr)

// Act
	_, err := mockBookHandlers.GetBooks()

// Assert
	if err == nil {
		t.Error("Test failed while returning an error")
	}
}
