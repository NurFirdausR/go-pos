package authentication

import (
	"context"
	"database/sql"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/repository/mysql/authentication"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUsecase struct {
	AuthenticateRepo authentication.AuthenticateRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewAuthenticationUsecase(authenticateRepo authentication.AuthenticateRepository, DB *sql.DB, validate *validator.Validate) UseCase {
	return &AuthenticateUsecase{
		AuthenticateRepo: authenticateRepo,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *AuthenticateUsecase) LoginHandler(ctx context.Context, reqLogin domain.User) (res domain.User) {
	err := service.Validate.Struct(reqLogin)
	helper.PanicIfError(err)
	// fmt.Println(service)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.AuthenticateRepo.GetByName(ctx, tx, reqLogin.Username)

	if err != nil {
		panic(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqLogin.Password))
	if err != nil {
		panic("Password Tidak Sesuai")
	}
	// helper.PanicIfError(err)
	res = domain.User{
		Id:       user.Id,
		Username: user.Username,
	}

	return res
}

func (service *AuthenticateUsecase) RegisterHandler(ctx context.Context, newUser domain.User) (res domain.User) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.AuthenticateRepo.GetByName(ctx, tx, newUser.Username)
	if err == nil {
		panic("Username Already Exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}
	newUser.Password = string(hashedPassword)
	res = service.AuthenticateRepo.CreateUser(ctx, tx, newUser)
	helper.PanicIfError(err)
	return res
}
