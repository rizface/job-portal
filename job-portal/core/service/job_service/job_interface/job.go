package job_service_interface

import (
	"job-portal/app/model/request"
	"job-portal/app/model/response"
)

type ManipulationJob interface {
	PostJob(companyName string, request request.Job) string
	DetailJob(jobId string, level string) response.Job
	DeleteJob(companyName,jobId string) string
	UpdateJob(request request.Job, companyName,jobId string) string
	TmpTakeDown(companyName,jobId string) string
}

type ApplicantFindJob interface {
	GetJobs(filter request.Filter) []response.Job
	Recommendation() []response.Job
}