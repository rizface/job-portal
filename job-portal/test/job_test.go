package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/app/setup"
	"job-portal/core/repository/job_repo"
	"job-portal/helper"
	filter_builder "job-portal/helper/filter"
	"job-portal/route"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	job := setup.FindJob()
	t.Run("get job", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet,"/jobs",nil)
			request.Header.Add("Authorization","Bearer " + generateTokenApplicant())
			recorder := httptest.NewRecorder()
			job.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("invalid token", func(t *testing.T) {
				request := httptest.NewRequest(http.MethodGet,route.JOBS,nil)
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusUnauthorized,recorder.Code)
			})

			t.Run("invalid paging", func(t *testing.T) {
				request := httptest.NewRequest(http.MethodGet,route.JOBS+"?start=sana",nil)
				request.Header.Add("Authorization","Bearer " + generateTokenApplicant())
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)	
				assert.Equal(t, http.StatusNotFound,recorder.Code)
			})
		})
	})
	
	t.Run("post job", func(t *testing.T) {
		router := setup.JobManipulationForCompany()
		t.Run("success", func(t *testing.T) {
			job := request.Job{
				Title:            "DevOps Engineer",
				Location:         "Padang, Indonesia",
				MinSalary:        5000000,
				MaxSalary:        10000000,
				Type:             "Intern",
				JobDescription:   "Melakukan Deploymen Service Yang Telah Dibuat",
				MinQualification: request.MinQualification{
					Requirements: []string{"have an alaytical thinking"},
					Prefered:     []string{"good looking"},
				},
				Status:           true,
				CreatedAt:        time.Now(),
			}
			payload,_ := json.Marshal(job)
			request := httptest.NewRequest(http.MethodPost,route.JOBS,bytes.NewReader(payload))
			request.Header.Add("Authorization", "Bearer " + generateCompanyToken())
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("invalid token", func(t *testing.T) {
				job := request.Job{}
				payload,_ := json.Marshal(job)
				request := httptest.NewRequest(http.MethodPost,route.JOBS,bytes.NewReader(payload))
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusUnauthorized,recorder.Code)
			})
			t.Run("empty request", func(t *testing.T) {
				job := request.Job{}
				payload,_ := json.Marshal(job)
				request := httptest.NewRequest(http.MethodPost,route.JOBS,bytes.NewReader(payload))
				request.Header.Add("Authorization", "Bearer " + generateCompanyToken())
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusBadRequest,recorder.Code)
			})
			job := request.Job{}
			t.Run("use applicant token", func(t *testing.T) {
				payload,_ := json.Marshal(job)
				request := httptest.NewRequest(http.MethodPost,route.JOBS,bytes.NewReader(payload))
				request.Header.Add("Authorization", "Bearer " + generateTokenApplicant())
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusForbidden,recorder.Code)
			})
		})
	})

	t.Run("detail job", func(t *testing.T) {
		repo := job_repo.NewFindJob()
		db := helper.Connection()
		ctx := context.Background()
		job := setup.DetailJob()
		t.Run("success", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet,"/jobs",nil)
			var jobs []response.Job
			result,_ := repo.GetJob(db,ctx,*filter_builder.BuildFilter(request))
			result.All(ctx,&jobs)
			id := jobs[0].Id
			request = httptest.NewRequest(http.MethodGet,"/jobs/"+id,nil)
			request.Header.Add("Authorization", "Bearer "+generateCompanyToken())
			recorder := httptest.NewRecorder()
			job.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("not found", func(t *testing.T) {
				request := httptest.NewRequest(http.MethodGet,"/jobs",nil)
				var jobs []response.Job
				result,_ := repo.GetJob(db,ctx,*filter_builder.BuildFilter(request))
				result.All(ctx,&jobs)
				id := jobs[0].Id
				request = httptest.NewRequest(http.MethodGet,"/jobs/"+id+"asd",nil)
				request.Header.Add("Authorization", "Bearer "+generateCompanyToken())
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusNotFound,recorder.Code)
			})
			t.Run("invalid token", func(t *testing.T) {
				request := httptest.NewRequest(http.MethodGet,"/jobs",nil)
				var jobs []response.Job
				result,_ := repo.GetJob(db,ctx,*filter_builder.BuildFilter(request))
				result.All(ctx,&jobs)
				id := jobs[0].Id
				request = httptest.NewRequest(http.MethodGet,"/jobs/"+id+"asd",nil)
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusUnauthorized,recorder.Code)
			})
		})
	})

	t.Run("update job", func(t *testing.T) {
		job := setup.JobManipulationForCompany()
		t.Run("success", func(t *testing.T) {
			payload := request.Job{
				Title:            "Senior Lead Devops Engineer",
				Location:         "Jakarta Selatan, Indonesia",
				MinSalary:        30000000,
				MaxSalary:        50000000,
				Type:             "Full Time",
				JobDescription:   "you will lead exprienced team to be the best devops team in the world",
			}
			payloadJson,_ := json.Marshal(payload)
			request := httptest.NewRequest(http.MethodPut,"/jobs/61dbbbe4740e57774c8391e9",bytes.NewReader(payloadJson))
			request.Header.Add("Authorization","Bearer " + generateCompanyToken())
			recorder := httptest.NewRecorder()
			job.ServeHTTP(recorder,request)
			assert.Equal(t, http.StatusOK,recorder.Code)
		})
		t.Run("failed", func(t *testing.T) {
			t.Run("invalid token", func(t *testing.T) {
				payload := request.Job{}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPut,"/jobs/61dbbbe4740e57774c8391e9",bytes.NewReader(payloadJson))
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusUnauthorized,recorder.Code)
			})
			t.Run("bad request", func(t *testing.T) {
				payload := request.Job{
				}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPut,"/jobs/61dbbbe4740e57774c8391e9",bytes.NewReader(payloadJson))
				request.Header.Add("Authorization","Bearer " + generateCompanyToken())
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusBadRequest,recorder.Code)

			})
			t.Run("not found (id is invalid)", func(t *testing.T) {
				payload := request.Job{
					Title:            "Senior Lead Devops Engineer",
					Location:         "Jakarta Selatan, Indonesia",
					MinSalary:        30000000,
					MaxSalary:        50000000,
					Type:             "Full Time",
					JobDescription:   "you will lead exprienced team to be the best devops team in the world",
				}
				payloadJson,_ := json.Marshal(payload)
				request := httptest.NewRequest(http.MethodPut,"/jobs/61d12sdbbe4740e57774c8391a1",bytes.NewReader(payloadJson))
				request.Header.Add("Authorization","Bearer " + generateCompanyToken())
				recorder := httptest.NewRecorder()
				job.ServeHTTP(recorder,request)
				assert.Equal(t, http.StatusNotFound,recorder.Code)

			})
		})
	})

}
