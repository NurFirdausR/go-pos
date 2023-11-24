package main

import (
	"net/http"

	authentication_controller "github.com/NurFirdausR/go-pos/controller/http/authentication"
	product_controller "github.com/NurFirdausR/go-pos/controller/http/product"
	"github.com/NurFirdausR/go-pos/database"
	"github.com/NurFirdausR/go-pos/exception"
	"github.com/NurFirdausR/go-pos/helper"
	authentication_repo "github.com/NurFirdausR/go-pos/repository/mysql/authentication"
	product_repo "github.com/NurFirdausR/go-pos/repository/mysql/product"
	authentication_usecase "github.com/NurFirdausR/go-pos/usecase/authentication"
	product_usecase "github.com/NurFirdausR/go-pos/usecase/product"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := database.Connect()

	router := httprouter.New()

	userRepo := authentication_repo.NewAuthenticationRepository()                          // Initialize repositories
	authUseCase := authentication_usecase.NewAuthenticationUsecase(userRepo, db, validate) // Initialize use cases
	authController := authentication_controller.NewAuthenticationController(authUseCase)   // Initialize handlers

	productRepo := product_repo.NewProductRepository()                             // Initialize repositories
	productUsecase := product_usecase.NewProductUsecase(productRepo, db, validate) // Initialize use cases
	productController := product_controller.NewProductController(productUsecase)   // Initialize handlers

	router.POST("/api/login", authController.LoginHandler)
	router.POST("/api/register", authController.RegisterHandler)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Save)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
