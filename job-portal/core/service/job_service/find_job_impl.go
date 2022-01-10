package job_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/core/repository/job_repo/job_interface"
	job_service_interface "job-portal/core/service/job_service/job_interface"
	"job-portal/helper"
	"time"
)

type findJob struct {
	db *mongo.Database
	valid *validator.Validate
	repo job_interface.ApplicantFindJob
}

func NewFindJob(db *mongo.Database, valid *validator.Validate, repo job_interface.ApplicantFindJob) job_service_interface.ApplicantFindJob {
	return &findJob{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (f *findJob) GetJobs(filter request.Filter) []response.Job {
	var result []response.Job
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	cursor,err := f.repo.GetJob(f.db,ctx,filter)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	err = cursor.All(ctx,&result)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	return result
}

func (f *findJob) Recommendation() []response.Job {
	panic("implement me")
}

