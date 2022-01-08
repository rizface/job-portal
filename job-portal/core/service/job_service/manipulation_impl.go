package job_service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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

func (j *jobManipulationService) PostJob(companyId string,request request.Job) string {
	valid := primitive.IsValidObjectID(companyId)
	helper.PanicException(exception.NotFouund{Err: "company tidak ditemukan"}, valid == false)
	err := j.valid.Struct(request)
	if err != nil {
		validation.Validation(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.PostJob(j.db,ctx,companyId,request)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "pekerjaan berhasil diposting"
	}
	return "pekerjaan gagal diposting"
}

func (j *jobManipulationService) DetailJob(jobId string) response.Job {
	var obj bson.M
	var result response.Job
	valid := primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	cursor := j.repo.DetailJob(j.db,ctx,jobId)
	err := cursor.Decode(&obj)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, errors.Is(err,mongo.ErrNoDocuments))
	bsonBytes,_ := bson.Marshal(obj["jobs"].(bson.A)[0])
	bson.Unmarshal(bsonBytes,&result)
	return result
}

func (j *jobManipulationService) DeleteJob(companyId,jobId string) string {
	valid := primitive.IsValidObjectID(companyId)
	helper.PanicException(exception.NotFouund{Err:"company tidak ditemukan"},valid == false)
	valid = primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"},valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.DeleteJob(j.db,ctx,companyId,jobId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "pekerjaan berhasil dihapus"
	}
	return "pekerjaan gagal dihapus"
}

func (j *jobManipulationService) UpdateJob(request request.Job, companyId,jobId string) string {
	err := j.valid.Struct(request)
	if err != nil {
		validation.Validation(err)
	}
	valid := primitive.IsValidObjectID(companyId)
	helper.PanicException(exception.NotFouund{Err:"company tidak ditemukan"}, valid == false)
	valid = primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := j.repo.UpdateJob(j.db,ctx,request,companyId,jobId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "pekerjaan berhasil diupdate"
	}
	return "pekerjaan gagal diupdate"
}

func (j *jobManipulationService) TmpTakeDown(companyId,jobId string) string {
	valid := primitive.IsValidObjectID(companyId)
	helper.PanicException(exception.NotFouund{Err:"company tidak ditemukan"}, valid == false)
	valid = primitive.IsValidObjectID(jobId)
	helper.PanicException(exception.NotFouund{Err:"pekerjaan tidak ditemukan"}, valid == false)
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	current := j.DetailJob(jobId)
	success,err := j.repo.TmpTakeDown(j.db,ctx,current,companyId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	if success {
		return "status pekerjaan berhasil diupdate"
	}
	return "status pekerjaan gagal diupdate"
}


