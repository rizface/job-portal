package job_interface

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
)

type ApplicantFindJon interface {
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
	DeleteJOb(db *mongo.Database, ctx context.Context, jobId string) (bool, error)
	UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, jobId string) (bool, error)
	TmpTakeDown(db *mongo.Database, ctx context.Context, current response.Job) (bool, error)
}
