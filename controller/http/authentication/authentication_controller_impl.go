package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/model/web"
	"github.com/NurFirdausR/go-pos/usecase/authentication"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		helper.PanicIfError(err)
	}
	result := cntrler.authenticate.LoginHandler(r.Context(), u)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Successfully!",
		Data:   result,
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

// func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	var u User
// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		helper.PanicIfError(err)
// 	}
// 	fmt.Printf("The request body is %v\n", u.Username)

// 	// Validate u using validator
// 	validate := validator.New()
// 	if err := validate.Struct(u); err != nil {
// 		helper.PanicIfError(err)
// 	}
// 	var count int
// 	// fmt.Printf(input.Username)
// 	query := "SELECT username,password FROM users WHERE username = ?"
// 	err := AuthService.DB.QueryRow(query, u.Username).Scan(&count)
// 	if err != nil {
// 		helper.PanicIfError(err)

// 	}
// 	if count > 0 {
// 		panic("Username Already Exists")
// 	}

// }

// func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// Parse and validate the request body
// 	var input User
// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		helper.PanicIfError(err)
// 	}

// 	// Validate input using validator
// 	validate := validator.New()
// 	if err := validate.Struct(input); err != nil {
// 		helper.PanicIfError(err)
// 	}

// 	var count int
// 	// fmt.Printf(input.Username)
// 	query := "SELECT COUNT(*) FROM users WHERE username = ?"
// 	err := AuthService.DB.QueryRow(query, input.Username).Scan(&count)
// 	if err != nil {
// 		helper.PanicIfError(err)

// 	}
// 	if count > 0 {
// 		panic("Username Already Exists")
// 	}

// 	// Hash the password using bcrypt
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		panic("Failed to hash password")
// 	}
// 	query = "INSERT INTO users (username,password) VALUES(?,?)"
// 	result, err := AuthService.DB.Exec(query, input.Username, string(hashedPassword))
// 	if err != nil {
// 		helper.PanicIfError(err)
// 	}
// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "Successfully!",
// 		Data:   result,
// 	}
// 	helper.WriteToResponseBody(w, webResponse)

// }
