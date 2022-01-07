package setup

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/core/controller/auth_controller"
	"job-portal/core/controller/auth_controller/applicant_auth_controller"
	"job-portal/core/controller/auth_controller/company_auth_controller"
	"job-portal/core/repository/auth_repo"
	"job-portal/core/service/auth_service/applicant_auth_service"
	"job-portal/core/service/auth_service/company_auth_service"
	"job-portal/helper"
)

var Db *mongo.Database
var Valid = validator.New()

func init() {
	//helper.LoadConfig(".env")
	Db = helper.Connection()
	//Redis = helper.Redis()
}

func ApplicantAuthController() auth_controller.Auth{
	repo := auth_repo.NewAuth()
	service := applicant_auth_service.NewAuth(Db,Valid,repo)
	controller := applicant_auth_controller.NewAuth(service)
	return controller
}

func CompanyAuthController() auth_controller.Auth {
	repo := auth_repo.NewAuth()
	service := company_auth_service.NewAuth(Db,Valid,repo)
	controller := company_auth_controller.NewAuth(service)
	return controller
}
