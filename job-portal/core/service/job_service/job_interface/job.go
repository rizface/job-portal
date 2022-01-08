package job_service_interface

import (
	"job-portal/app/model/request"
	"job-portal/app/model/response"
)

type ManipulationJob interface {
	PostJob(companyId string, request request.Job) string
	DetailJob(jobId string) response.Job
	DeleteJob(companyId,jobId string) string
	UpdateJob(request request.Job, companyId,jobId string) string
	TmpTakeDown(companyId,jobId string) string
}
