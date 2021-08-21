package bookRepository

import (
	"books-list/domain"
	"books-list/err"
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
		appErr.Message = "unexpected database error"
		return nil, &appErr
	}
	return books, nil
}

func (b BookRepositoryDb) GetBook(id int) (*domain.Book, *err.Error) {
	sqlRequest := "select * from books_list where id=$1"
	err := b.db.Get(&book, sqlRequest, id)
	if err != nil {
		if err == sql.ErrNoRows {
			appErr.Message = "Book ID Not Found"
			return nil, &appErr
		} else {
			appErr.Message = "Server error"
			return nil, &appErr
		}
	}
	return &book, nil
}

func (b BookRepositoryDb) AddBook(book domain.Book) (int, error) {
	var id int
	stmt, err := b.db.PrepareNamed("insert into books_list (title, author, year) values(:title, :author, :year) RETURNING id;")
	if err != nil {
		return 0, err
	}
	err = stmt.Get(&id, book)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b BookRepositoryDb) UpdateBook(book domain.Book) (int64, error) {
	result, err := b.db.Exec("update books_list set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		return 0, err
	}
	// ошибка при выведении RowsAffected
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

func (b BookRepositoryDb) RemoveBook(id int) (int64, error) {
	result, err := b.db.Exec("delete from books_list where id = $1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, _ := result.RowsAffected()
	return rowsDeleted, nil
}

func NewBookRepositoryDb(db *sqlx.DB) BookRepositoryDb {
	return BookRepositoryDb{
		db,
	}
}
