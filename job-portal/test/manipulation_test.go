package test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/core/repository/job_repo"
	"job-portal/helper"
	"testing"
)

func TestPostJob(t *testing.T) {
	t.Run("repository", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			repo := job_repo.NewManipulationJob()
			repo.PostJob(helper.Connection(), context.Background(), "61d7f0a6400ef82c291d5876", request.Job{
				Title:  "junior backend",
				Detail: "ini kerja untuk junior backend",
			})
		})
	})
}

func TestDeleteJob(t *testing.T) {
	t.Run("repository", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			repo := job_repo.NewManipulationJob()
			repo.DeleteJOb(helper.Connection(), context.Background(), "61d84e096b9ff3da2a297483")
		})
	})
}

func TestUpdateStatusJobs(t *testing.T) {
	db := helper.Connection()
	t.Run("repository", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			var result bson.M
			var res response.Job
			repo := job_repo.NewManipulationJob()
			current := repo.DetailJob(db, context.Background(), "61d83dcde8c45531c61695bb")
			current.Decode(&result)
			bsonBytes, _ := bson.Marshal(result["jobs"].(bson.A)[0])
			bson.Unmarshal(bsonBytes, &res)
			repo.TmpTakeDown(db, context.Background(), res)
		})
	})
}

func TestUpdateJobs(t *testing.T) {
	db := helper.Connection()
	t.Run("repository", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			repo := job_repo.NewManipulationJob()
			repo.UpdateJob(db, context.Background(), request.Job{
				Title:     "junior backend",
				Detail:    "ini kerja untuk junior backend",
				MinSalary: 5000000,
				MaxSalary: 100000000,
			},"61d84261163b46b2f3d4d92a")
		})
	})
}
