package applicant_auth_service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/app/validation"
	myredis "job-portal/core/redis"
	"job-portal/core/repository/auth_repo"
	"job-portal/core/service/auth_service"
	"job-portal/helper"
	"time"
)

type auth struct {
	db *mongo.Database
	valid *validator.Validate
	repo auth_repo.Auth
}

func NewAuth(db *mongo.Database, valid *validator.Validate,repo auth_repo.Auth) auth_service.Auth {
	return &auth{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (a *auth) Login(request request.Auth,collection string) string {
	var result response.Applicant
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	request.Level = "applicant"
	cursor := a.repo.Login(a.db,ctx,request,collection)
	err := cursor.Decode(&result)
	helper.PanicException(exception.NotFound{Err: "akun kamu tidak terdaftar"},errors.Is(err,mongo.ErrNoDocuments))
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"},err != nil)
	helper.PanicException(exception.Forbidden{Err: "lakukan aktifasi akun terlebih dahulu,cek email kamu"}, result.Status == false)
	err = helper.ComparePassword(request.Password,result.Password)
	helper.PanicException(exception.BadRequest{Err:"email / password kamu salah"},err != nil)
	token,err := helper.GenerateApplicantToken(result)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"},err != nil)
	return token
}

func (a *auth) Register(request request.Auth,collection string) string {
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	err := a.valid.Struct(request)
	if err != nil {
		validation.Validation(err)
	}
	exists := a.repo.Login(a.db,ctx,request,"applicants")
	helper.PanicException(exception.Conflict{Err:"akun ini sudah terdaftar"}, exists.Err() == nil)

	hash,err := helper.GeneratePassword(request.Password)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	request.Password = hash
	request.Level = "applicant"
	defer cancel()
	insertId,err := a.repo.Register(a.db,ctx,request,collection)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	activateUrl := "applicant/"+insertId.(string)
	myredis.RedisPublish("job-portal-email",request.Email,activateUrl)
	return "registrasi berhasil, periksa email kamu untuk melakukan aktifasi akun"
}


func (a *auth) Aktivasi(userId string) bool {
	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	result := a.repo.Aktivasi(a.db,ctx,userId,"applicants")
	return result
}