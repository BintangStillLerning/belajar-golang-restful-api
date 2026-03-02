package service

import (
	"bintangakasyah/belajar-golang-restful-api/exception"
	"bintangakasyah/belajar-golang-restful-api/helper"
	"bintangakasyah/belajar-golang-restful-api/model/domain"
	"bintangakasyah/belajar-golang-restful-api/model/web"
	"bintangakasyah/belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	db                 *sql.DB
	Validator *validator.Validate

}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		db:                 db,
		Validator: validator,
	}
}
	

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, err)

	category := domain.Category{
		Name: request.Name,
	}

	savedCategory := service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(savedCategory)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)
	
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, err)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
    if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	updatedCategory := service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(updatedCategory)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, err)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, err)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	 if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, err)

	categories, err := service.CategoryRepository.FindAll(ctx, tx)
    helper.ToCategoryResponses(categories)
	return []web.CategoryResponse{}
}