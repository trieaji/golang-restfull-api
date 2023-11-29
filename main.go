package main

import (
	"golang-restfullapi/app"
	"golang-restfullapi/controller"
	"golang-restfullapi/helper"
	"golang-restfullapi/middleware"
	"golang-restfullapi/repository"
	"golang-restfullapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB() 
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server {
		Addr: "localhost:4002",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}