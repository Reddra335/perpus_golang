package service

import (
	"context"
	"database/sql"
	"perpus_golang/exception"
	"perpus_golang/helper"
	"perpus_golang/model/domain"
	"perpus_golang/model/web"
	"perpus_golang/repository"

	"github.com/go-playground/validator"
)

type BookServiceImpln struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewBookServiceImpln(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) *BookServiceImpln {

	return &BookServiceImpln{
		BookRepository: bookRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *BookServiceImpln) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.validate.Struct(request)
	helper.ErrorT(err)
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	helper.CommitOrRollback(tx)

	commitData := domain.Book{
		Title:      request.Title,
		Author:     request.Author,
		Isbn:       request.Isbn,
		Stock:      request.Stock,
		CategoryId: request.CategoryId,
		FilePath:   request.FilePath,
	}
	result := service.BookRepository.Save(ctx, tx, commitData)
	return helper.ToBookResponse(result)

}

func (service *BookServiceImpln) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.validate.Struct(request)
	helper.ErrorT(err)
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	helper.CommitOrRollback(tx)

	result, errFindId := service.BookRepository.FindById(ctx, tx, request.Id)
	if errFindId != nil {
		exception.NewNotFound(errFindId.Error())
	}

	resultUpdate := service.BookRepository.Update(ctx, tx, result)

	return helper.ToBookResponse(resultUpdate)
}

func (service *BookServiceImpln) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	helper.CommitOrRollback(tx)

	FindId, errFindId := service.BookRepository.FindById(ctx, tx, bookId)
	if errFindId != nil {

		exception.NewNotFound(errFindId.Error())
	}

	service.BookRepository.Delete(ctx, tx, FindId)

}

func (service *BookServiceImpln) FindById(ctx context.Context, bookId int) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	helper.CommitOrRollback(tx)

	result, errFindById := service.BookRepository.FindById(ctx, tx, bookId)
	if errFindById != nil {
		exception.NewNotFound(errFindById.Error())
	}

	return helper.ToBookResponse(result)

}

func (service *BookServiceImpln) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	helper.CommitOrRollback(tx)

	result := service.BookRepository.FindAll(ctx, tx)
	return helper.ToBookResponseSlice(result)

}
