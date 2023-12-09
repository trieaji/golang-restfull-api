package helper

import (
	"golang-restfullapi/model/domain"
	"golang-restfullapi/model/web"
)

//untuk findById
func ToCategoryResponse(category domain.Category) web.CategoryResponse {//function ini hanya berfungsi untuk mengembalikan response
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

//untuk findAll
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {//function ini hanya berfungsi untuk mengembalikan response
	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}

	return categoryResponse
}