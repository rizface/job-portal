package company_profile_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/exception"
	"job-portal/app/model"
	"job-portal/app/model/response"
	"job-portal/app/validation"
	"job-portal/core/repository/profile_repo"
	"job-portal/core/service/profile_service"
	"job-portal/helper"
	"time"
)

type profile struct {
	db *mongo.Database
	valid *validator.Validate
	repo profile_repo.BasicProfile
}

func NewProfile(db *mongo.Database, valid *validator.Validate, repo profile_repo.BasicProfile) profile_service.BasicProfile {
	return &profile{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}


func (p *profile) GetDetail(userId string) model.Object {
	company := response.NewCompany()
	valid := primitive.IsValidObjectID(userId)
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	cursor := p.repo.GetProfile(ctx,p.db,userId)
	err := cursor.Decode(company)
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, err != nil)
	return company

}

func (p *profile) UpdateDetail(userId string, user model.Object) string {
	valid := primitive.IsValidObjectID(userId)
	err := p.valid.Struct(user)
	if err != nil {
		validation.Validation(err)
	}
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	bsonUser := user.Convert()
	success,err := p.repo.UpdateProfile(ctx,p.db,userId,bsonUser)
	helper.PanicException(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"}, err != nil)
	if success == false {
		return "akun gagal diupdate"
	}
	return "akun berhasil diupdate"
}

