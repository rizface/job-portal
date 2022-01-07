package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/exception"
	"job-portal/app/model/response"
	"os"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

type Applicant struct {
	Id                             primitive.ObjectID
	NamaDepan, NamaBelakang, Email string
	Status                         bool
	jwt.RegisteredClaims
	_ struct{}
}

type Company struct {
	Id                    primitive.ObjectID
	NamaPerusahaan, Email string
	Status                bool
	jwt.RegisteredClaims
	_ struct{}
}

func GenerateApplicantToken(applicant response.Applicant) (string, error) {
	claim := Applicant{
		Id:           applicant.Id,
		NamaDepan:    applicant.NamDepan,
		NamaBelakang: applicant.NamaBelakang,
		Email:        applicant.Email,
		Status:       applicant.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func GenerateCompanyToken(company response.Company) (string, error) {
	claim := Company{
		Id:             company.Id,
		NamaPerusahaan: company.NamaPerusahaan,
		Email:          company.Email,
		Status:         company.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyApplicantToken(tokenString string) interface{} {
	claims := &Applicant{}
	tkn,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret),nil
	})
	fmt.Println(tkn.Valid)
	PanicException(exception.UnAuthorized{Err:"token tidak valid"},tkn.Valid == false || err != nil && errors.Is(err,jwt.ErrSignatureInvalid))
	PanicException(exception.UnAuthorized{Err:"terjadi kesalahan pada sistem kami"},err != nil )
	return claims
}

func VerifyCompanyToken(tokenString string) interface{} {
	claims := &Company{}
	tkn,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret),nil
	})

	PanicException(exception.UnAuthorized{Err:"token tidak valid"},tkn.Valid == false || err != nil && errors.Is(err,jwt.ErrSignatureInvalid))
	PanicException(exception.UnAuthorized{Err:"terjadi kesalahan pada sistem kami"},err != nil )
	return claims
}