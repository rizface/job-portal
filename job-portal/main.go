package main

import (
	"job-portal/app/config"
	exception2 "job-portal/app/exception"
	"job-portal/app/setup"
	"job-portal/helper"
	"net/http"
)

func main() {
	setup.Auth()
	setup.ApplicantProfile()
	setup.CompanyProfile()
	setup.JobManipulationForCompany()
	setup.DetailJob()
	setup.FindJob()

	err := http.ListenAndServe(":"+config.APP_PORT,setup.Mux)
	helper.PanicException(exception2.InternalServerError{Err:"terjadi kesalahan saat memulai server"},err != nil)
}
