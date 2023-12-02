package exception

import (
	"fmt"
	"net/http"

	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if validationError(w, r, err) {
		return
	}
	if unauthorizedError(w, r, err) {
		return
	}

	if noCookie(w, r, err) {
		return
	}
	if jwtErrSignature(w, r, err) {
		return
	}
	if jwtErrTokenExp(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		errors := map[string]interface{}{}
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
		}

		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errors,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "Aplication/Json")
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func noCookie(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if err == http.ErrNoCookie {
		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}
func jwtErrSignature(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if err == jwt.ErrSignatureInvalid {
		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func jwtErrTokenExp(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if err == jwt.ErrTokenExpired {
		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized, Token Expired",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	if err == http.StatusUnauthorized {
		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized, Token Invalid",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}
