package setup

import 	(
	"github.com/gorilla/mux"
	"job-portal/app/middleware"
	"job-portal/route"
	"net/http"
)

var Mux = mux.NewRouter()

func Auth() *mux.Router {
	applicantController := ApplicantAuthController()
	companyController := CompanyAuthController()
	router := Mux.NewRoute().Subrouter()
	router.Use(middleware.ErrorHandler)

	router.HandleFunc(route.LOGIN_APPLICANT,applicantController.Login).Methods(http.MethodPost)
	router.HandleFunc(route.REGISTER_APPLICANT,applicantController.Register).Methods(http.MethodPost)
	router.HandleFunc(route.ACTIVATE_APPLICANT,applicantController.Aktivasi).Methods(http.MethodGet)


	router.HandleFunc(route.LOGIN_COMPANY,companyController.Login).Methods(http.MethodPost)
	router.HandleFunc(route.REGISTER_COMPANY,companyController.Register).Methods(http.MethodPost)
	router.HandleFunc(route.ACTIVATE_COMPANY,companyController.Aktivasi).Methods(http.MethodGet)
	return router
}

func ApplicantProfile() *mux.Router {
	applicant := ApplicantProfileController()
	router := Mux.NewRoute().Subrouter()
	router.Use(middleware.ErrorHandler,middleware.ApplicantAuth)
	router.HandleFunc(route.APPLICANT_PROFILE,applicant.GetDetail).Methods(http.MethodGet)
	router.HandleFunc(route.APPLICANT_PROFILE,applicant.UpdateDetail).Methods(http.MethodPut)
	return router
}

func CompanyProfile() *mux.Router {
	company := CompanyProfileController()
	router := Mux.NewRoute().Subrouter()
	router.Use(middleware.ErrorHandler,middleware.CompanyAuth)
	router.HandleFunc(route.COMPANY_PROFILE,company.GetDetail).Methods(http.MethodGet)
	router.HandleFunc(route.COMPANY_PROFILE,company.UpdateDetail).Methods(http.MethodPut)
	return router
}

func JobManipulationForCompany() *mux.Router {
	controller := JobManipulationController()
	router := Mux.NewRoute().Subrouter()
	router.Use(middleware.ErrorHandler,middleware.CompanyAuth,middleware.CompanyCompleteProfile)
	router.HandleFunc(route.JOB_MANIPULATION,controller.UpdateJob).Methods(http.MethodPut)
	router.HandleFunc(route.JOB_MANIPULATION,controller.DetailJob).Methods(http.MethodGet)
	router.HandleFunc(route.JOB_MANIPULATION,controller.DeleteJob).Methods(http.MethodDelete)
	router.HandleFunc(route.JOBS,controller.PostJob).Methods(http.MethodPost)
	router.HandleFunc(route.JOB_TAKEDOWN,controller.TakeDown).Methods(http.MethodPut)
	return router
}