package helper

import (
	"perpus_golang/model/domain"
	"perpus_golang/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:         book.CategoryId,
		Title:      book.Title,
		Author:     book.Author,
		Isbn:       book.Isbn,
		Stock:      book.Stock,
		CategoryId: book.CategoryId,
		FilePath:   book.FilePath,
	}
}
func ToCategoryResponseSlice(categories []domain.Category) []web.CategoryResponse {

	var ResponseCategory []web.CategoryResponse

	for _, category := range categories {
		ResponseCategory = append(ResponseCategory, ToCategoryResponse(category))

	}
	return ResponseCategory

}

func ToBookResponseSlice(books []domain.Book) []web.BookResponse {
	var ManyBook []web.BookResponse

	for _, book := range books {
		ManyBook = append(ManyBook, ToBookResponse(book))

	}
	return ManyBook
}
