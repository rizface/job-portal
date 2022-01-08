package job_interface

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
)

type ApplicantFindJob interface {
	GetJob()
	JobRecomendation()
}

type ApplicantApplication interface {
	Apply()
	Applied()
	CancelPropose()
	UpdatePropose()
}

type ManipulationJob interface {
	PostJob(db *mongo.Database, ctx context.Context, companyId string, request request.Job) (bool, error)
	DetailJob(db *mongo.Database, ctx context.Context, jobId string) *mongo.SingleResult
	DeleteJob(db *mongo.Database, ctx context.Context, companyId,jobId string) (bool, error)
	UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, companyId,jobId string) (bool, error)
	TmpTakeDown(db *mongo.Database, ctx context.Context, current response.Job,companyId string) (bool, error)
}
