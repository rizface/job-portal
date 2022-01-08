package job_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/app/validation"
	"job-portal/core/repository/job_repo/job_interface"
	job_service_interface "job-portal/core/service/job_service/job_interface"
	"job-portal/helper"
	"time"
)

type jobManipulationService struct {
	db *mongo.Database
	valid *validator.Validate
	repo job_interface.ManipulationJob
}

func NewJobManipulation(db *mongo.Database, valid *validator.Validate, repo job_interface.ManipulationJob) job_service_interface.ManipulationJob {
	return &jobManipulationService{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (j *jobManipulationService) PostJob(companyName string,request request.Job) string {
	err := j.valid.Struct(request)
	if err != nil {
		validation.Validation(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.PostJob(j.db,ctx,companyName,request)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success != nil {
		return "pekerjaan berhasil diposting"
	}
	return "pekerjaan gagal diposting"
}

func (j *jobManipulationService) DetailJob(jobId string) response.Job {
	var result response.Job
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	cursor := j.repo.DetailJob(j.db,ctx,jobId)
	err := cursor.Decode(&result)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, errors.Is(err,mongo.ErrNoDocuments))
	return result
}

func (j *jobManipulationService) DeleteJob(companyName,jobId string) string {
	valid := primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"},valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.DeleteJob(j.db,ctx,companyName,jobId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "pekerjaan berhasil dihapus"
	}
	return "pekerjaan gagal dihapus"
}

func (j *jobManipulationService) UpdateJob(request request.Job, companyName,jobId string) string {
	err := j.valid.Struct(request)
	if err != nil {
		validation.Validation(err)
	}
	valid := primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.UpdateJob(j.db,ctx,request,companyName,jobId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "pekerjaan berhasil diupdate"
	}
	return "pekerjaan gagal diupdate"
}

func (j *jobManipulationService) TmpTakeDown(companyName,jobId string) string {
	valid := primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	current := j.DetailJob(jobId)
	fmt.Println(current)
	success,err := j.repo.TmpTakeDown(j.db,ctx,current,companyName)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "status pekerjaan berhasil diupdate"
	}
	return "status pekerjaan gagal diupdate"
}


