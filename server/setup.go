package server

import (
	"challenges_8/config"
	"challenges_8/module/controller"
	repository "challenges_8/module/repository/book"
	service "challenges_8/module/service/book"
)

type Ctrs struct {
	BookCtr controller.BookController
}

func InitControllers() Ctrs {
	//why err if i do this?
	dataStore := config.NewDBPostges()

	bookRepo := repository.NewBookRepositoryImpl(dataStore)
	bookServ := service.NewBookServiceImpl(bookRepo)
	bookCtr := controller.NewBookControllerImpl(bookServ)

	return Ctrs{
		BookCtr: bookCtr,
	}
}
