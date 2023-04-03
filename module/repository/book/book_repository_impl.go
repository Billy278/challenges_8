package repository

import (
	model "challenges_8/module/model/book"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepositoryImpl(dt *sql.DB) BookRepository {
	return &BookRepositoryImpl{
		DB: dt,
	}
}

func (bookRepo *BookRepositoryImpl) FindByIdBook(ctx context.Context, idIn uint64) (book model.Book, err error) {
	sql := "SELECT id,title,author,des,create_at,update_at FROM books  WHERE id=$1 AND delete_at IS NULL"
	row, err := bookRepo.DB.QueryContext(ctx, sql, idIn)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.Create_at, &book.Update_at)
		if err != nil {
			return
		}
		return
	} else {
		return book, errors.New("NOT FOUND")
	}
}
func (bookRepo *BookRepositoryImpl) CreateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	sql := "INSERT INTO books(title,author,des,create_at) VALUES($1,$2,$3,$4) RETURNING id,title,author,des,create_at"
	rows, err := bookRepo.DB.QueryContext(ctx, sql, bookIn.Title, bookIn.Author, bookIn.Desc, bookIn.Create_at)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.Create_at)
		if err != nil {
			return
		}
	}
	return
}
func (bookRepo *BookRepositoryImpl) UpdateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	sql := "UPDATE books SET title=$1,author=$2,des=$3,update_at=$4 WHERE id=$5 RETURNING id,title,author,des,update_at"
	rows, err := bookRepo.DB.QueryContext(ctx, sql, bookIn.Title, bookIn.Author, bookIn.Desc, bookIn.Update_at, bookIn.Id)
	if err != nil {
		return
	}
	if rows.Next() {
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.Update_at)
		if err != nil {
			return
		}
		return
	} else {
		return book, errors.New("FAILED UPDATE BOOK")
	}

}

func (bookRepo *BookRepositoryImpl) FindAllBook(ctx context.Context) (books []model.Book, err error) {
	sql := "SELECT id,title,author,des,create_at,update_at FROM books WHERE delete_at IS NULL"
	rows, err := bookRepo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}

	for rows.Next() {
		book := model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.Create_at, &book.Update_at)
		if err != nil {
			return
		}
		books = append(books, book)
	}
	return
}
func (bookRepo *BookRepositoryImpl) SoftDeleteBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {

	sql := "UPDATE books SET delete_at=$1 WHERE id=$2 RETURNING id"
	rows, err := bookRepo.DB.QueryContext(ctx, sql, bookIn.Delete_at, bookIn.Id)
	if err != nil {
		return
	}
	if rows.Next() {
		rows.Scan(&book.Id)
	}
	return
}

func (bookRepo *BookRepositoryImpl) HardDeleteBook(ctx context.Context, bookId uint64) (book model.Book, err error) {
	sql := "DELETE FROM books WHERE id=$1"
	_, err = bookRepo.DB.ExecContext(ctx, sql, bookId)
	if err != nil {
		return
	}
	book.Id = bookId
	return
}

func (bookRepo *BookRepositoryImpl) FindByIdBookSoftDelete(ctx context.Context, idIn uint64) (book model.Book, err error) {
	sql := "SELECT id,title,author,des,create_at,update_at FROM books  WHERE id=$1"
	row, err := bookRepo.DB.QueryContext(ctx, sql, idIn)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.Create_at, &book.Update_at)
		if err != nil {
			return
		}
		return
	} else {
		return book, errors.New("NOT FOUND")
	}
}
