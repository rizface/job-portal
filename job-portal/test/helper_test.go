package test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/helper"
	"math/rand"
	"strconv"
	"time"
)

func generateAuthRequest() request.Auth {
	rand.Seed(time.Now().Unix())
	email := strconv.Itoa(rand.Intn(1000))+"@gmail.com"
	password,_ := helper.GeneratePassword("rahasia")
	return request.Auth{
		Email: email,
		Password: password,
	}
}

func aktivasi(email string, level string) {
	db := helper.Connection()
	db.Collection(level).UpdateOne(context.Background(),bson.M{"email":email},bson.M{"$set":bson.M{"status":true}})
}

func generateTokenApplicant() string {
	token,_ := helper.GenerateApplicantToken(response.Applicant{
		Id:           primitive.NewObjectID(),
		NamDepan:     "saya",
		NamaBelakang: "kamu",
		Password:     "rahasia",
		Email:        "rahasia@gmail.com",
		Level:        "applicant",
		Status:       true,
	})
	return token
}

func generateCompanyToken() string {
	token,_ := helper.GenerateCompanyToken(response.Company{
		Id:               primitive.NewObjectID(),
		Email:            "company@gmail.com",
		NamaPerusahaan:   "company",
		DetailPerusahaan: "company keren",
		LinkWebsite:      "company.com",
		Alamat:           "bumi",
		Level:            "company",
		Password:         "rahasia",
		Status:           true,
	})
	return token
}
