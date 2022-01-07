package company_auth_controller

import (
	"encoding/json"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/core/controller/auth_controller"
	"job-portal/core/service/auth_service"
	"job-portal/helper"
	"net/http"
)

type auth struct {
	service auth_service.Auth
}

func NewAuth(service auth_service.Auth) auth_controller.Auth {
	return &auth{service: service}
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	var data request.Auth
	err := json.NewDecoder(r.Body).Decode(&data)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kamu"},err != nil)
	token := a.service.Login(data,"companies")
	helper.JsonWriter(w,http.StatusOK,"login berhasil",map[string]interface{} {
		"token":token,
	})
}

func (a *auth) Register(w http.ResponseWriter, r *http.Request) {
	var data request.Auth
	err := json.NewDecoder(r.Body).Decode(&data)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kamu"},err != nil)
	result := a.service.Register(data,"companies")
	helper.JsonWriter(w,http.StatusOK,result,nil)
}


func (a *auth) Aktivasi(w http.ResponseWriter, r *http.Request) {
	result := a.service.Aktivasi(helper.GetParamsValue(r,"userId"))
	if result == true {
		helper.JsonWriter(w,http.StatusOK,"aktivasi berhasil,silahkan login untuk melanjutkan",nil)
	} else {
		helper.JsonWriter(w,http.StatusOK,"aktivasi gagal",nil)
	}
}
