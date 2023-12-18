package category

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
