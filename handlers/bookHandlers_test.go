package handlers

import (
	domain2 "books-list/domain"
	"books-list/mocks/repository"
	"strconv"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockBookHandlers *domain.MockBookRepository
var router *mux.Router
var ctrl *gomock.Controller
var bh BookHandlers

func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockBookHandlers = domain.NewMockBookRepository(ctrl)
	router = mux.NewRouter()
	bh = BookHandlers{Repository: mockBookHandlers}
}

func Test_GetBooks_should_return_error (t *testing.T) {
	setup(t)
// Arrange
	mockBookHandlers.EXPECT().GetBooks().Return(nil, &error)
	router.HandleFunc("/books", bh.GetBooks)
	request, _ := http.NewRequest(http.MethodGet, "/books", nil)

// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("error-error-error")
	}
}

func Test_GetBook_should_return_error (t *testing.T) {
	setup(t)
// Arrange
	router.HandleFunc("/books/1001", bh.GetBook)
	request, _ := http.NewRequest(http.MethodGet, "/books/1001", nil)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	mockBookHandlers.EXPECT().GetBook(id).Return(nil, &error)

// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

// Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("error-error-error")
	}
}

func Test_should_return_books (t *testing.T) {
	setup(t)
	dummyBooks := []domain2.Book{
				{	ID:     13,
					Title:  "Pavel",
					Author: "book",
					Year:   "2021",
				},
				{	ID: 14,
					Author: "Lena",
					Title: "books",
					Year: "2020",
				},
			}
	mockBookHandlers.EXPECT().GetBooks().Return(dummyBooks, nil)
	_, err := mockBookHandlers.GetBooks()
	if err != nil {
		t.Error("Test failed")
	}
}