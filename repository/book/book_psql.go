package bookRepository

import (
	"books-list/models"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sqlx.DB, books []models.Book) ([]models.Book, error) {

	sqlRequest := "select * from books_list"
	err := db.Select(&books, sqlRequest)
	if err != nil {
		return []models.Book{}, err
	}
	return books, nil
}

func (b BookRepository) GetBook(db *sqlx.DB, book models.Book, id int) (models.Book, error) {
	sqlRequest := "select * from books_list where id=$1"

	err := db.Get(&book, sqlRequest, id)
	if err != nil {
		return models.Book{}, err
	}
	return book, err
}

func (b BookRepository) AddBook(db *sqlx.DB, book models.Book) (int, error) {
	stmt, err := db.PrepareNamed("insert into books_list (title, author, year) values(:title, :author, :year) RETURNING id;")
	if err != nil {
		return 0, err
	}
	var id int
	err = stmt.Get(&id, book)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b BookRepository) UpdateBook(db *sqlx.DB, book models.Book) (int64, error) {
	result, err := db.Exec("update books_list set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sqlx.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books_list where id = $1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}
