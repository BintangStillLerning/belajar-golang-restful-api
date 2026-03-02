package main

import (
	"bintangakasyah/belajar-golang-restful-api/app"
	"bintangakasyah/belajar-golang-restful-api/controller"
	"bintangakasyah/belajar-golang-restful-api/exception"
	"bintangakasyah/belajar-golang-restful-api/helper"
	"bintangakasyah/belajar-golang-restful-api/middleware"
	"bintangakasyah/belajar-golang-restful-api/repository"
	"bintangakasyah/belajar-golang-restful-api/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)


func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router:= httprouter.New()
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:6767",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()

	port := os.Getenv("PORT")
	if port == ""{
		port = "6767"
	}
	http.ListenAndServe(":" + port, middleware.NewAuthMiddleware(router))
	helper.PanicIfError(err)
}