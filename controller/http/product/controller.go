package product

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
