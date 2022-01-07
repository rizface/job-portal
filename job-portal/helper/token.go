package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/model/response"
	"os"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

type Applicant struct {
	Id                                       primitive.ObjectID
	NamaDepan, NamaBelakang, Username, Email string
	Status                                   bool
	jwt.RegisteredClaims
	_ struct{}
}

type Company struct {
	Id                              primitive.ObjectID
	NamaPerusahaan, Username, Email string
	Status                          bool
	jwt.RegisteredClaims
	_ struct{}
}

func GenerateApplicantToken(applicant response.Applicant) (string, error) {
	claim := Applicant{
		Id:           primitive.ObjectID{},
		NamaDepan:    applicant.NamDepan,
		NamaBelakang: applicant.NamaBelakang,
		Username:     applicant.Username,
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
		Id:       primitive.ObjectID{},
		NamaPerusahaan: company.NamaPerusahaan,
		Email:    company.Email,
		Username: company.Username,
		Status:   company.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
