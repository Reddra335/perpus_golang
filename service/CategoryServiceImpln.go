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

type CategoryServiceImpln struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryServiceImpln(CategoryService repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpln {

	return &CategoryServiceImpln{
		CategoryRepository: CategoryService,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpln) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.ErrorT(err)
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}
func (service *CategoryServiceImpln) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.ErrorT(err)
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpln) Delete(ctx context.Context, CategoryId int) {
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, CategoryId)

	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpln) FindById(ctx context.Context, CategoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, CategoryId)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpln) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := service.DB.Begin()
	helper.ErrorT(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponseSlice(categories)
}
