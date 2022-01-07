package setup

import (
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
