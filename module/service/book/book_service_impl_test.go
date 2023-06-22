package service

import (
	model "challenges_8/module/model/book"
	repoMock "challenges_8/module/repository/book/mock"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	createAt := time.Now()
	type (
		input struct {
			bookIn model.Book
		}
		want struct {
			Bookres []model.Book
			err     error
		}
	)
	testCases := []struct {
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		desc   string
		doMock func(repoMock *repoMock.MockBookRepository)
		input  input
		want   want
	}{
		{
			desc: "Happy case",
			input: input{
				bookIn: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test Author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			want: want{
				Bookres: []model.Book{{
					Id:        1,
					Title:     "test title",
					Author:    "test Author",
					Desc:      "test desc",
					Create_at: &createAt,
				}},
				err: nil,
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().FindAllBook(gomock.Any()).Return(
					[]model.Book{
						{
							Id:        1,
							Title:     "test title",
							Author:    "test Author",
							Desc:      "test desc",
							Create_at: &createAt,
						},
					}, nil).MaxTimes(1).AnyTimes()
			},
		},
		{
			desc: "Fail case",
			input: input{
				bookIn: model.Book{
					Title:     "test title",
					Author:    "test Author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			want: want{
				Bookres: []model.Book{{}},
				err:     errors.New("Error"),
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().FindAllBook(gomock.Any()).Return(
					[]model.Book{}, errors.New("Error"),
				).MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := repoMock.NewMockBookRepository(ctrl)
			tC.doMock(repoMock)
			svc := BookServiceImpl{
				RepoBook: repoMock,
			}
			findall, err := svc.FindAllBookSrv(context.Background())
			if err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.want.Bookres, findall)
			}

		})
	}
}

func TestCreateBook(t *testing.T) {
	createAt := time.Now()
	type (
		input struct {
			bookIn model.Book
		}
		want struct {
			Bookres model.Book
			err     error
		}
	)
	testCases := []struct {
		desc   string
		doMock func(repoMock *repoMock.MockBookRepository)
		input  input
		want   want
	}{
		{
			desc: "Happy case",
			input: input{
				bookIn: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			want: want{
				Bookres: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(
					model.Book{
						Id:        1,
						Title:     "test title",
						Author:    "test author",
						Desc:      "test desc",
						Create_at: &createAt,
					}, nil,
				).MaxTimes(1).AnyTimes()
			},
		},
		{
			desc: "Fail case",
			input: input{
				bookIn: model.Book{
					Title:     "test title",
					Author:    "test author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			want: want{
				Bookres: model.Book{},
				err:     errors.New("Bad Request"),
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(
					model.Book{}, errors.New("Bad Request"),
				).MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := repoMock.NewMockBookRepository(ctrl)
			tC.doMock(repoMock)
			svc := BookServiceImpl{
				RepoBook: repoMock,
			}
			res, err := svc.CreateBookSrv(context.Background(), tC.input.bookIn)
			if err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.want.Bookres, res)
			}

		})
	}

}

func TestFindByid(t *testing.T) {
	createAt := time.Now()
	type (
		input struct {
			bookId uint64
		}
		want struct {
			Bookres model.Book
			err     error
		}
	)
	testCases := []struct {
		desc   string
		input  input
		want   want
		doMock func(repoMock *repoMock.MockBookRepository)
	}{
		{
			desc: "Happy case",
			input: input{
				bookId: uint64(1),
			},
			want: want{
				Bookres: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test author",
					Desc:      "test desc",
					Create_at: &createAt,
				},
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().FindByIdBook(gomock.Any(), gomock.Any()).Return(
					model.Book{
						Id:        1,
						Title:     "test title",
						Author:    "test author",
						Desc:      "test desc",
						Create_at: &createAt,
					}, nil,
				).MaxTimes(1).AnyTimes()
			},
		},
		{
			desc: "Fail result",
			input: input{
				bookId: uint64(1),
			},
			want: want{
				Bookres: model.Book{},
				err:     errors.New("NOT FOUND"),
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().FindByIdBook(gomock.Any(), gomock.Any()).Return(
					model.Book{}, errors.New("NOT FOUND"),
				).MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := repoMock.NewMockBookRepository(ctrl)
			tC.doMock(repoMock)
			svcBook := BookServiceImpl{
				RepoBook: repoMock,
			}
			res, err := svcBook.FindByIdBookSrv(context.Background(), tC.input.bookId)
			if err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, res, tC.want.Bookres)
			}
		})
	}

}

func TestUpdateBook(t *testing.T) {
	createAt := time.Now()
	updateAt := time.Now()
	type (
		input struct {
			bookIn model.Book
		}
		want struct {
			Bookres model.Book
			err     error
		}
	)
	testCases := []struct {
		desc   string
		input  input
		want   want
		doMock func(repoMock *repoMock.MockBookRepository)
	}{
		{
			desc: "Happy case",
			input: input{
				bookIn: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test auth",
					Desc:      "test des",
					Create_at: &createAt,
				},
			},
			want: want{
				Bookres: model.Book{
					Id:        1,
					Title:     "test title",
					Author:    "test auth",
					Desc:      "test des",
					Create_at: &createAt,
					Update_at: &updateAt,
				},
			},
			doMock: func(repoMock *repoMock.MockBookRepository) {
				repoMock.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(
					model.Book{
						Id:        1,
						Title:     "test title",
						Author:    "test auth",
						Desc:      "test des",
						Create_at: &createAt,
						Update_at: &updateAt,
					}, nil,
				).MaxTimes(1).AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := repoMock.NewMockBookRepository(ctrl)
			srvBooks := BookServiceImpl{
				RepoBook: repoMock,
			}
			res, _ := srvBooks.UpdateBookSrv(context.Background(), tC.input.bookIn)
			assert.Equal(t, tC.want.Bookres, res)
		})
	}

}
