package bookRepository

import (
	"books-list/domain"
	"books-list/dto"
	"github.com/jmoiron/sqlx"
)

// implementation of the Secondary Port
type BookRepositoryDb struct {
	db *sqlx.DB
}

func (b BookRepositoryDb) GetBooks(books []domain.Book) ([]domain.Book, error) {
	sqlRequest := "select * from books_list"
	err := b.db.Select(&books, sqlRequest)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b BookRepositoryDb) GetBook(book domain.Book, id int) (domain.Book, error) {
	sqlRequest := "select * from books_list where id=$1"

	err := b.db.Get(&book, sqlRequest, id)
	if err != nil {
		return domain.Book{}, err
	}
	return book, err
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

func (b BookRepositoryDb) UpdateBook(updateBookRequest dto.UpdateBookRequest) (int64, error) {
	book := domain.Book{
		ID: updateBookRequest.Id,
		Title: updateBookRequest.Title,
		Author: updateBookRequest.Author,
		Year: updateBookRequest.Year,
	}
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
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}

func NewBookRepositoryDb(db *sqlx.DB) BookRepositoryDb {
	return BookRepositoryDb{
		db,
	}
}
