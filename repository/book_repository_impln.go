package repository

import (
	"context"
	"database/sql"
	"errors"
	"perpus_golang/helper"
	"perpus_golang/model/domain"
)

type BookRepositoryImpln struct{}

func NewBookRepositoryImpln() BookRepository {
	return &BookRepositoryImpln{}
}

func (repository *BookRepositoryImpln) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "INSERT INTO book(title, author, isbn, stock, category_id, file_path) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, book.Title, book.Author, book.Isbn, book.Stock, book.CategoryId, book.FilePath)
	helper.ErrorT(err)

	id, err := result.LastInsertId()
	helper.ErrorT(err)

	book.Id = int(id)

	return book

}
func (repository *BookRepositoryImpln) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "UPDATE book SET title = ?, author = ?, isbn = ?, stock = ?, category_id = ?, file_path = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Title, book.Author, book.Isbn, book.Stock, book.CategoryId, book.FilePath, book.Id)
	helper.ErrorT(err)

	return book

}
func (repository *BookRepositoryImpln) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	SQL := "DELETE FROM book WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.ErrorT(err)

}
func (repository *BookRepositoryImpln) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	SQL := "SELECT id, title,author,isbn,stock,category_id,file_path FROM book WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.ErrorT(err)

	defer rows.Close()
	book := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Stock, &book.CategoryId, book.FilePath)
		helper.ErrorT(err)

		return book, nil
	} else {

		return book, errors.New("Data Not Found")
	}

}
func (repository *BookRepositoryImpln) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "SELECT id, title, author, isbn, stock, category_id, file_path FROM book"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.ErrorT(err)
	defer rows.Close()

	books := []domain.Book{}
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Stock, &book.CategoryId, book.FilePath)
		helper.ErrorT(err)

		books = append(books, book)
	}
	return books
}
