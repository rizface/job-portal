package job_controller

import (
	"encoding/json"
	"job-portal/app/exception"
	request2 "job-portal/app/model/request"
	"job-portal/core/controller/job_controller/job_controller_interface"
	job_service_interface "job-portal/core/service/job_service/job_interface"
	"job-portal/helper"
	"net/http"
)

type jobManipulation struct {
	service job_service_interface.ManipulationJob
}

func NewJobManipulation(service job_service_interface.ManipulationJob)  job_controller_interface.Manipulation{
	return &jobManipulation{service: service}
}

func (j *jobManipulation) PostJob(w http.ResponseWriter, r *http.Request) {
	var request request2.Job
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	companyId := helper.GetTokenValue(r.Context().Value("company-data"),"NamaPerusahaan")
	result := j.service.PostJob(companyId,request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (j *jobManipulation) DetailJob(w http.ResponseWriter, r *http.Request) {
	jobId := helper.GetParamsValue(r,"jobId")
	result := j.service.DetailJob(jobId)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{}{
		"job":result,
	})
}

func (j *jobManipulation) DeleteJob(w http.ResponseWriter, r *http.Request) {
	companyId := helper.GetTokenValue(r.Context().Value("company-data"),"NamaPerusahaan")
	jobId := helper.GetParamsValue(r,"jobId")
	result := j.service.DeleteJob(companyId,jobId)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (j *jobManipulation) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var request request2.Job
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	companyId := helper.GetTokenValue(r.Context().Value("company-data"),"NamaPerusahaan")
	jobId := helper.GetParamsValue(r,"jobId")
	result := j.service.UpdateJob(request,companyId,jobId)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (j *jobManipulation) TakeDown(w http.ResponseWriter, r *http.Request) {
	companyId := helper.GetTokenValue(r.Context().Value("company-data"),"NamaPerusahaan")
	jobId := helper.GetParamsValue(r,"jobId")
	result := j.service.TmpTakeDown(companyId,jobId)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}


