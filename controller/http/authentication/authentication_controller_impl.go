package authentication

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NurFirdausR/go-pos/config"
	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/NurFirdausR/go-pos/usecase/authentication"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// type RegistrationInput struct {
// 	Username string `json:"username" validate:"required,max=20"`
// 	Password string `json:"password" validate:"required,max=20"`
// }

type authenticateController struct {
	authenticate authentication.UseCase
}

func NewAuthenticationController(authenticate authentication.UseCase) AuthController {
	return &authenticateController{
		authenticate: authenticate,
	}

}

func (cntrler *authenticateController) LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		helper.PanicIfError(err)
	}
	// validasi data
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		helper.PanicIfError(err)
	}

	// menjalankan function login
	result := cntrler.authenticate.LoginHandler(r.Context(), u)
	// pembuatan JWT Token
	expTime := time.Now().Add(time.Minute * 60)
	claims := config.JWTClaim{
		Username: result.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-pos",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	//deklarasikan algoritma untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)

	//set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Successfully!",
		// Data:   },
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (cntrler *authenticateController) RegisterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var u domain.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		helper.PanicIfError(err)
	}
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		helper.PanicIfError(err)
	}
	result := cntrler.authenticate.RegisterHandler(r.Context(), u)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Successfully!",
		Data:   result,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (cntrler *authenticateController) LogoutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Hapus token di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Successfully!",
		// Data:   },
	}
	helper.WriteToResponseBody(w, webResponse)
}
