package main

import (
	"net/http"

	authentication_controller "github.com/NurFirdausR/go-pos/controller/http/authentication"
	"github.com/NurFirdausR/go-pos/database"
	"github.com/NurFirdausR/go-pos/exception"
	"github.com/NurFirdausR/go-pos/helper"
	authentication_repo "github.com/NurFirdausR/go-pos/repository/mysql/authentication"
	authentication_usecase "github.com/NurFirdausR/go-pos/usecase/authentication"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := database.Connect()

	router := httprouter.New()

	// Initialize repositories
	userRepo := authentication_repo.NewAuthenticationRepository()

	// Initialize use cases
	authUseCase := authentication_usecase.NewAuthenticationUsecase(userRepo, db, validate)

	// Initialize handlers
	authController := authentication_controller.NewAuthenticationController(authUseCase)

	router.POST("/api/login", authController.LoginHandler)
	router.POST("/api/register", authController.RegisterHandler)
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
