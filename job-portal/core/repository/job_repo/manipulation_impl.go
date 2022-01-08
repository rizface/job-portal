package job_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
	"job-portal/app/model/response"
	"job-portal/core/repository/job_repo/job_interface"
)

type manipulationJob struct{}

func NewManipulationJob() job_interface.ManipulationJob {
	return &manipulationJob{}
}

func (m manipulationJob) PostJob(db *mongo.Database, ctx context.Context, companyName string, request request.Job) (interface{}, error) {
	data := request.Convert()
	data["company"] = companyName
	result,err := db.Collection("jobs").InsertOne(ctx,data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (m manipulationJob) DetailJob(db *mongo.Database, ctx context.Context, jobId string) *mongo.SingleResult {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	cursor := db.Collection("jobs").FindOne(ctx,bson.M{"_id":objJob})
	return cursor
}

func (m manipulationJob) DeleteJob(db *mongo.Database, ctx context.Context, companyName,jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	result,err := db.Collection("jobs").DeleteMany(ctx,bson.M{
		"$and":bson.A{
			bson.M{"company":companyName},
			bson.M{"_id":objJob},
		},
	})
	if err != nil {
		return false,err
	}
	return result.DeletedCount > 0, nil
}

func (m manipulationJob) UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, companyName,jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	result,err := db.Collection("jobs").UpdateOne(ctx,bson.M{
		"$and":bson.A{
			bson.M{"_id":objJob},
			bson.M{"company":companyName},
		},
	},bson.M{
		"$set": bson.M{
			"title":      request.Title,
			"detail":     request.Detail,
			"min_salary": request.MinSalary,
			"max_salary": request.MaxSalary,
		},
	})
	if err != nil {
		return false,nil
	}
	return result.ModifiedCount > 0, err
}

func (m manipulationJob) TmpTakeDown(db *mongo.Database, ctx context.Context,current response.Job,companyName string) (bool, error) {
	result,err := db.Collection("jobs").UpdateOne(ctx,bson.M{
		"$and":bson.A{
			bson.M{"company":companyName},
			bson.M{"_id":current.Id},
		},
	},bson.M{
		"$set":bson.M{
			"status":!current.Status,
		},
	})
	if err != nil{
		return false,err
	}
	return result.ModifiedCount > 0,nil
}
