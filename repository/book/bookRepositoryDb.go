package bookRepository

import (
	"books-list/domain"
	"books-list/err"
	"books-list/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// implementation of the Secondary Port
type BookRepositoryDb struct {
	db *sqlx.DB
}

var appErr err.Error
var books []domain.Book
var book domain.Book

func (b BookRepositoryDb) GetBooks() ([]domain.Book, *err.Error) {
	sqlRequest := "select * from books_list"
	err := b.db.Select(&books, sqlRequest)
	if err != nil {
		logger.Error("Database error while getting all books")
		appErr.Message = "Unexpected server error"
		return nil, &appErr
	}
	return books, nil
}

func (b BookRepositoryDb) GetBook(id int) (*domain.Book, *err.Error) {
	sqlRequest := "select * from books_list where id=$1"
	err := b.db.Get(&book, sqlRequest, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while getting a book from a database: book ID not found")
			appErr.Message = "Book ID Not Found"
			return nil, &appErr
		} else {
			logger.Error("Unexpected database error")
			appErr.Message = "Server error"
			return nil, &appErr
		}
	}
	return &book, nil
}

func (b BookRepositoryDb) AddBook(book domain.Book) (int, *err.Error) {
	var id int
	stmt, err := b.db.PrepareNamed("insert into books_list (title, author, year) values(:title, :author, :year) RETURNING id;")
	if err != nil {
		logger.Error("Error while making SQL request in AddBook function")
		appErr.Message = "Unexpected server error"
		return 0, &appErr
	}
	err = stmt.Get(&id, book)
	if err != nil {
		logger.Error("Error while parsing last ID number")
		appErr.Message = "Unexpected server error"
		return 0, &appErr
	}
	return id, nil
}

func (b BookRepositoryDb) UpdateBook(book domain.Book) (int, *err.Error) {
	result, err := b.db.Exec("update books_list set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		logger.Error("Error while making SQL request in UpdateBook function")
		appErr.Message = "Unexpected server error"
		return 0, &appErr
	}
	// ошибка при выведении RowsAffected
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error while calculating rows affected in UpdateBook function")
		appErr.Message = "Unexpected server error"
		return 0, &appErr
	}
	return int(rowsUpdated), nil
}

func (b BookRepositoryDb) RemoveBook(id int) (int, *err.Error) {
	result, err := b.db.Exec("delete from books_list where id = $1", id)
	if err != nil {
		logger.Error("Error while executing a delete SQL query")
		appErr.Message = "Server error while deleting a book"
		return 0, &appErr
	}
	rowsDeleted, _ := result.RowsAffected()
	return int(rowsDeleted), nil
}

func NewBookRepositoryDb(db *sqlx.DB) BookRepositoryDb {
	return BookRepositoryDb{
		db,
	}
}