package repository

import (
	model "challenges_8/module/model/book"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepoImplMongo struct {
	DB *mongo.Database
}

func NewBookRepositoryImplMongo(db *mongo.Database) BookRepository {
	return &BookRepoImplMongo{
		DB: db,
	}
}

func (bookrepo *BookRepoImplMongo) CreateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	_, err = bookrepo.DB.Collection("data_buku").InsertOne(ctx, bookIn)
	if err != nil {
		return
	}
	return
}

func (bookrepo *BookRepoImplMongo) UpdateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	filter := bson.M{
		"id": bookIn.Id,
	}
	_, err = bookrepo.DB.Collection("data_buku").UpdateOne(ctx, filter, bson.M{
		"$set": bookIn,
	})
	if err != nil {
		return
	}
	return
}
func (bookrepo *BookRepoImplMongo) FindByIdBook(ctx context.Context, idIn uint64) (book model.Book, err error) {
	filter := bson.M{
		"id": idIn,
	}

	err = bookrepo.DB.Collection("data_buku").FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return
	}
	return
}
func (bookrepo *BookRepoImplMongo) FindAllBook(ctx context.Context) (books []model.Book, err error) {
	cur, err := bookrepo.DB.Collection("data_buku").Find(ctx, bson.D{})
	if err != nil {
		return
	}
	book := model.Book{}
	for cur.Next(ctx) {
		err = cur.Decode(&book)
		if err != nil {
			return
		}
		books = append(books, book)
	}
	return
}
func (bookrepo *BookRepoImplMongo) HardDeleteBook(ctx context.Context, bookId uint64) (book model.Book, err error) {
	filter := bson.M{
		"id": bookId,
	}
	_, err = bookrepo.DB.Collection("data_buku").DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	return
}
func (bookrepo *BookRepoImplMongo) SoftDeleteBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	return
}

func (bookrepo *BookRepoImplMongo) FindByIdBookSoftDelete(ctx context.Context, idIn uint64) (book model.Book, err error) {
	return
}
