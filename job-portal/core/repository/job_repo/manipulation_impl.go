package job_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	data["_id"] = primitive.NewObjectID()
	data["company"] = companyName
	result, err := db.Collection("jobs").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (m manipulationJob) DetailJob(db *mongo.Database, ctx context.Context, jobId string, isCompany bool) *mongo.SingleResult {
	var cursor *mongo.SingleResult
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	if isCompany {
		cursor = db.Collection("jobs").FindOne(ctx, bson.M{"_id": objJob})
	} else {
		cursor = db.Collection("jobs").FindOne(ctx,bson.M{"_id":objJob},options.FindOne().SetProjection(bson.M{
			"applicants":0,
		}))
	}
	return cursor
}

func (m manipulationJob) DeleteJob(db *mongo.Database, ctx context.Context, companyName, jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	result, err := db.Collection("jobs").DeleteMany(ctx, bson.M{
		"$and": bson.A{
			bson.M{"company": companyName},
			bson.M{"_id": objJob},
		},
	})
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}

func (m manipulationJob) UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, companyName, jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	result, err := db.Collection("jobs").UpdateOne(ctx, bson.M{
		"$and": bson.A{
			bson.M{"_id": objJob},
			bson.M{"company": companyName},
		},
	}, bson.M{
		"$set": bson.M{
			"title":             request.Title,
			"location":          request.Location,
			"min_salary":        request.MinSalary,
			"max_salary":        request.MaxSalary,
			"type":              request.Type,
			"job_description":   request.JobDescription,
			"min_qualification": request.MinQualification,
		},
	})
	if err != nil {
		return false, nil
	}
	return result.ModifiedCount > 0, err
}

func (m manipulationJob) TmpTakeDown(db *mongo.Database, ctx context.Context, current response.Job, companyName string) (bool, error) {
	objId,_ := primitive.ObjectIDFromHex(current.Id)
	result, err := db.Collection("jobs").UpdateOne(ctx, bson.M{
		"$and": bson.A{
			bson.M{"company": companyName},
			bson.M{"_id": objId},
		},
	}, bson.M{
		"$set": bson.M{
			"status": !current.Status,
		},
	})
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}
