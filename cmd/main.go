package main

import (
	"net/http"

	authentication_controller "github.com/NurFirdausR/go-pos/controller/http/authentication"
	category_controller "github.com/NurFirdausR/go-pos/controller/http/category"
	product_controller "github.com/NurFirdausR/go-pos/controller/http/product"
	"github.com/NurFirdausR/go-pos/database"
	"github.com/NurFirdausR/go-pos/exception"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/middleware"
	authentication_repo "github.com/NurFirdausR/go-pos/repository/mysql/authentication"
	category_repo "github.com/NurFirdausR/go-pos/repository/mysql/category"
	product_repo "github.com/NurFirdausR/go-pos/repository/mysql/product"
	authentication_usecase "github.com/NurFirdausR/go-pos/usecase/authentication"
	category_usecase "github.com/NurFirdausR/go-pos/usecase/category"
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

	categoryRepo := category_repo.NewCategoryRepository()                              // Initialize repositories
	categoryUsecase := category_usecase.NewCategoryUsecase(categoryRepo, db, validate) // Initialize use cases
	categoryController := category_controller.NewCategoryController(categoryUsecase)   // Initialize handlers

	router.POST("/login", authController.LoginHandler)
	router.POST("/register", authController.RegisterHandler)
	router.GET("/logout", authController.LogoutHandler)

	router.GET("/api/products", middleware.JWTMiddleware(productController.FindAll))
	router.GET("/api/products/:productId", middleware.JWTMiddleware(productController.FindById))
	router.POST("/api/products", middleware.JWTMiddleware(productController.Save))
	router.PUT("/api/products/:productId", middleware.JWTMiddleware(productController.Update))
	router.DELETE("/api/products/:productId", middleware.JWTMiddleware(productController.Delete))

	router.POST("/api/categories", middleware.JWTMiddleware(categoryController.Save))

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
