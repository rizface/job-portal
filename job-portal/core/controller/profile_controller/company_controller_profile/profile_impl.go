package company_controller_profile

import (
	"encoding/json"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/core/controller/profile_controller"
	"job-portal/core/service/profile_service"
	"job-portal/helper"
	"net/http"
)

type profile struct {
	service profile_service.BasicProfile
}

func NewProfile(service profile_service.BasicProfile) profile_controller.BasicProfile {
	return &profile{
		service: service,
	}
}

func (p *profile) GetDetail(w http.ResponseWriter, r *http.Request) {
	userId := helper.GetTokenValue(r.Context().Value("company-data"),"id")
	result := p.service.GetDetail(userId)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"data":result,
	})
}

func (p *profile) UpdateDetail(w http.ResponseWriter, r *http.Request) {
	company := request.NewCompany()
	userId := helper.GetTokenValue(r.Context().Value("company-data"),"id")
	err := json.NewDecoder(r.Body).Decode(company)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	result := p.service.UpdateDetail(userId, company)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

