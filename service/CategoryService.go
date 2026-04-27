package service

import (
	"context"
	"perpus_golang/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.WebResponse
	SaveAll(ctx context.Context) []web.WebResponse
}
