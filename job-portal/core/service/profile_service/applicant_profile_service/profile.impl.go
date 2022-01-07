package applicant_profile_service

import (
	"context"
	"fmt"
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
	applicant := response.NewApplicant()
	valid := primitive.IsValidObjectID(userId)
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	cursor := p.repo.GetProfile(ctx,p.db,userId)
	err := cursor.Decode(applicant)
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, err != nil)
	return applicant

}

func (p *profile) UpdateDetail(userId string, user model.Object) string {
	valid := primitive.IsValidObjectID(userId)
	helper.PanicException(exception.NotFouund{Err:"akun tidak ditemukan"}, valid == false)
	err := p.valid.Struct(user)
	if err != nil {
		validation.Validation(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	bsonUser := user.Convert()
	success,err := p.repo.UpdateProfile(ctx,p.db,userId,bsonUser)
	fmt.Println(err)
	helper.PanicException(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"}, err != nil)
	if success == false {
		return "akun gagal diupdate"
	}
	return "akun berhasil diupdate"
}

