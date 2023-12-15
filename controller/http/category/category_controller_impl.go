package category

import (
	"encoding/json"
	"net/http"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/NurFirdausR/go-pos/usecase/category"

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
