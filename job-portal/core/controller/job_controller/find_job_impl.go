package job_controller

import (
	"job-portal/core/controller/job_controller/job_controller_interface"
	job_service_interface "job-portal/core/service/job_service/job_interface"
	"job-portal/helper"
	"job-portal/helper/filter"
	"net/http"
)

type findJob struct {
	service job_service_interface.ApplicantFindJob
}

func NewFindJob(service job_service_interface.ApplicantFindJob) job_controller_interface.ApplicantFindJob {
	return &findJob{service: service}
}

func (f *findJob) GetJobs(w http.ResponseWriter, r *http.Request) {
	filter := filter_builder.BuildFilter(r)
	result := f.service.GetJobs(*filter)
	status,link := filter_builder.GetPrevNext(len(result),filter.Skip,filter.Limit)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"jobs":result,
		status:link,
	})
}

func (f *findJob) Recomentdation(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

