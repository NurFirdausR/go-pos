package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/NurFirdausR/go-pos/usecase/product"
	"github.com/julienschmidt/httprouter"
)

type productControllerImpl struct {
	ProductUsecase product.UseCase
}

func NewProductController(productUsecase product.UseCase) ProductController {
	return &productControllerImpl{
		ProductUsecase: productUsecase,
	}

}

func (controller *productControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	productResponse, err := controller.ProductUsecase.FindById(r.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	productResponses := controller.ProductUsecase.FindAll(r.Context())
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   productResponses,
	}
	helper.WriteToResponseBody(w, WebResponse)
}

func (controller *productControllerImpl) Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Mendekode body permintaan HTTP menjadi struktur domain.Product
	decoder := json.NewDecoder(r.Body)
	productCreateRequst := domain.Product{}
	err := decoder.Decode(&productCreateRequst)
	helper.PanicIfError(err)

	// Menyimpan produk menggunakan use case yang sesuai
	productResponse := controller.ProductUsecase.Save(r.Context(), productCreateRequst)

	// Menyiapkan respon HTTP dengan data produk yang telah disimpan
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   productResponse,
	}

	// Menulis respon ke body HTTP
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	// Mendekode body permintaan HTTP menjadi struktur domain.Product
	decoder := json.NewDecoder(r.Body)
	productUpdateRequest := domain.Product{}
	err = decoder.Decode(&productUpdateRequest)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := controller.ProductUsecase.Update(r.Context(), productUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(w, webResponse)

}

func (controller *productControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productDeleteRequestt := domain.Product{}
	productDeleteRequestt.Id = id
	controller.ProductUsecase.Delete(r.Context(), productDeleteRequestt)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}
