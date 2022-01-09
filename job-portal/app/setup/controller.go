package setup

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/core/controller/auth_controller"
	"job-portal/core/controller/auth_controller/applicant_auth_controller"
	"job-portal/core/controller/auth_controller/company_auth_controller"
	"job-portal/core/controller/job_controller"
	"job-portal/core/controller/job_controller/job_controller_interface"
	"job-portal/core/controller/profile_controller"
	"job-portal/core/controller/profile_controller/applicant_controller_profile"
	"job-portal/core/controller/profile_controller/company_controller_profile"
	"job-portal/core/repository/auth_repo"
	"job-portal/core/repository/job_repo"
	"job-portal/core/repository/profile_repo/applicant_repo_profile"
	"job-portal/core/repository/profile_repo/company_repo_profile"
	"job-portal/core/service/auth_service/applicant_auth_service"
	"job-portal/core/service/auth_service/company_auth_service"
	"job-portal/core/service/job_service"
	"job-portal/core/service/profile_service/applicant_profile_service"
	"job-portal/core/service/profile_service/company_profile_service"
	"job-portal/helper"
)

var Db *mongo.Database
var Valid = validator.New()

func init() {
	//helper.LoadConfig(".env")
	Db = helper.Connection()
	//Redis = helper.Redis()
}

func applicantAuthController() auth_controller.Auth{
	repo := auth_repo.NewAuth()
	service := applicant_auth_service.NewAuth(Db,Valid,repo)
	controller := applicant_auth_controller.NewAuth(service)
	return controller
}

func companyAuthController() auth_controller.Auth {
	repo := auth_repo.NewAuth()
	service := company_auth_service.NewAuth(Db,Valid,repo)
	controller := company_auth_controller.NewAuth(service)
	return controller
}

func applicantProfileController() profile_controller.BasicProfile {
	repo := applicant_repo_profile.NewProfile()
	service := applicant_profile_service.NewProfile(Db,Valid,repo)
	controller := applicant_controller_profile.NewProfile(service)
	return controller
}

func companyProfileController() profile_controller.BasicProfile {
	repo := company_repo_profile.NewProfile()
	service := company_profile_service.NewProfile(Db,Valid,repo)
	controller := company_controller_profile.NewProfile(service)
	return controller
}

func jobManipulationController() job_controller_interface.Manipulation {
	repo := job_repo.NewManipulationJob()
	service := job_service.NewJobManipulation(Db,Valid,repo)
	controller := job_controller.NewJobManipulation(service)
	return controller
}

func findJobController() job_controller_interface.ApplicantFindJob {
	repo := job_repo.NewFindJob()
	service := job_service.NewFindJob(Db,Valid,repo)
	controller := job_controller.NewFindJob(service)
	return controller
}