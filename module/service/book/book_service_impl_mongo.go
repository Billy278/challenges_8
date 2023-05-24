package service

import (
	model "challenges_8/module/model/book"
	repository "challenges_8/module/repository/book"
	"context"
	"log"
	"time"
)

type BookServImplMongo struct {
	RepoBookMongo repository.BookRepository
}

func NewBookServImplMongo(repobook repository.BookRepository) BookService {
	return &BookServImplMongo{
		RepoBookMongo: repobook,
	}
}

func (booksrv_imp *BookServImplMongo) CreateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tNow := time.Now()
	bookIn.Create_at = &tNow
	book, err = booksrv_imp.RepoBookMongo.CreateBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error Insert Book :%v\n", err)
	}
	return
}
func (booksrv_imp *BookServImplMongo) UpdateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBookMongo.FindByIdBook(ctx, bookIn.Id)
	if err != nil {
		return book, err
	}
	book.Title = bookIn.Title
	book.Author = bookIn.Author
	book.Desc = bookIn.Desc
	book, err = booksrv_imp.RepoBookMongo.UpdateBook(ctx, book)
	if err != nil {
		log.Printf("[ERROR] error Update Book :%v\n", err)
	}
	return
}
func (booksrv_imp *BookServImplMongo) FindByIdBookSrv(ctx context.Context, idIn uint64) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBookMongo.FindByIdBook(ctx, idIn)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return
	}
	return
}
func (booksrv_imp *BookServImplMongo) FindAllBookSrv(ctx context.Context) (books []model.Book, err error) {
	books, err = booksrv_imp.RepoBookMongo.FindAllBook(ctx)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return
	}
	return
}
func (booksrv_imp *BookServImplMongo) SoftDeleteBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	return
}
func (booksrv_imp *BookServImplMongo) HardDeleteBookSrv(ctx context.Context, bookId uint64) (book model.Book, err error) {
	_, err = booksrv_imp.RepoBookMongo.FindByIdBook(ctx, bookId)
	if err != nil {
		return book, err
	}
	book, err = booksrv_imp.RepoBookMongo.HardDeleteBook(ctx, bookId)
	if err != nil {
		log.Printf("[ERROR] error deletebook Book :%v\n", err)
		return
	}
	return
}
func (booksrv_imp *BookServImplMongo) FindByIdBookSoftDeleteSrv(ctx context.Context, idIn uint64) (book model.Book, err error) {
	return
}
