package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/NurFirdausR/go-pos/usecase/category"
	category_web "github.com/NurFirdausR/go-pos/web/category"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryUsecase category.UseCase
}

func NewCategoryController(categoryUsecase category.UseCase) CategoryController {
	return &CategoryControllerImpl{
		CategoryUsecase: categoryUsecase,
	}
}

func (controller *CategoryControllerImpl) Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := domain.Category{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryUsecase.Save(r.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponse,
	}

	// Menulis respon ke body HTTP
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponses, err := controller.CategoryUsecase.FindById(r.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}
func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	decoder := json.NewDecoder(r.Body)
	categoryUpdateRequest := category_web.UpdateCategoryRequest{}
	err = decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryUsecase.Update(r.Context(), categoryUpdateRequest)
	fmt.Println("test")
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
