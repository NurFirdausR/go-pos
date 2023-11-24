package authentication

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	RegisterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
