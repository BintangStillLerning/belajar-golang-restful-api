package controller

import (
	"bintangakasyah/belajar-golang-restful-api/helper"
	"bintangakasyah/belajar-golang-restful-api/model/web"
	"bintangakasyah/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoyControllerImpl{
		CategoryService: categoryService,
	}
}

type CategoyControllerImpl struct {
 CategoryService service.CategoryService
}

func (controller *CategoyControllerImpl) Create(w http.ResponseWriter, request *http.Request, params httprouter.Params) {

	  categoryCreateRequest := web.CategoryCreateRequest{}
	  helper.ReadFromRequestBody(request, &categoryCreateRequest)

	  controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	  webResponse := web.WebResponse{
		 Code : 200,
		 Status: "OK",
		 Data: web.CategoryResponse{},
	  }
	  helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoyControllerImpl) Update(w http.ResponseWriter, request *http.Request, params httprouter.Params)  {

	  categoryUpdateRequest := web.CategoryUpdateRequest{}
      helper.ReadFromRequestBody(request, &categoryUpdateRequest)
	 categoryId := params.ByName("categoryId")
     id, err := strconv.Atoi(categoryId)
	 helper.PanicIfError(err)

	 categoryUpdateRequest.Id = id

	 categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	 webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data: categoryResponse,
	 }

	 helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoyControllerImpl) Delete(w http.ResponseWriter, request *http.Request, params httprouter.Params) {


	 categoryId := params.ByName("categoryId")
     id, err := strconv.Atoi(categoryId)
	 helper.PanicIfError(err)

	 controller.CategoryService.Delete(request.Context(), id)
	 webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
	 }

	 helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoyControllerImpl) FindById(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
       categoryId := params.ByName("categoryId")
     id, err := strconv.Atoi(categoryId)
	 helper.PanicIfError(err)

	 categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	 webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data: categoryResponse,
	 }

	 helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoyControllerImpl) FindAll(w http.ResponseWriter, request *http.Request, params httprouter.Params) {

	 categoryResponse := controller.CategoryService.FindAll(request.Context())
	 webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data: categoryResponse,
	 }
	 helper.WriteToResponseBody(w, webResponse)
}