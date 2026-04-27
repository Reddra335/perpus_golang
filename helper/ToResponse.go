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
func ToCategoryResponseSlice(categories []domain.Category) []web.CategoryResponse {

	var ResponseCategory []web.CategoryResponse

	for _, category := range categories {
		ResponseCategory = append(ResponseCategory, ToCategoryResponse(category))

	}
	return ResponseCategory

}
