package job_interface

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
)

type ApplicantFindJob interface {
	GetJob(db *mongo.Database, ctx context.Context, filter request.Filter) (*mongo.Cursor,error)
	JobRecomendation(db *mongo.Database, ctx context.Context, interset []string) (*mongo.Cursor,error)
}

type ApplicantApplication interface {
	Apply()
	Applied()
	CancelPropose()
	UpdatePropose()
}

type ManipulationJob interface {
	PostJob(db *mongo.Database, ctx context.Context, companyName string, request request.Job) (interface{}, error)
	DetailJob(db *mongo.Database, ctx context.Context, jobId string, isCompany bool) *mongo.SingleResult
	DeleteJob(db *mongo.Database, ctx context.Context, companyName,jobId string) (bool, error)
	UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, companyName,jobId string) (bool, error)
	TmpTakeDown(db *mongo.Database, ctx context.Context, current response.Job,companyName string) (bool, error)
}
