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

func (m manipulationJob) PostJob(db *mongo.Database, ctx context.Context, companyId string, request request.Job) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(companyId)
	result, err := db.Collection("companies").UpdateOne(ctx, bson.M{
		"_id": bson.M{"$eq": objId},
	}, bson.M{
		"$push": bson.M{
			"jobs": request.Convert(),
		},
	})
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (m manipulationJob) DetailJob(db *mongo.Database, ctx context.Context, jobId string) *mongo.SingleResult {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	cursor := db.Collection("companies").FindOne(ctx, bson.M{
		"jobs.id": objJob,
	},options.FindOne().SetProjection(bson.M{"jobs.$":1}))
	return cursor
}

func (m manipulationJob) DeleteJob(db *mongo.Database, ctx context.Context, companyId,jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	objCompany,_ := primitive.ObjectIDFromHex(companyId)
	result, err := db.Collection("companies").UpdateOne(ctx, bson.M{
		"_id":objCompany,
	}, bson.M{
		"$pull": bson.M{
			"jobs": bson.M{"id": objJob},
		},
	})
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (m manipulationJob) UpdateJob(db *mongo.Database, ctx context.Context, request request.Job, companyId,jobId string) (bool, error) {
	objJob, _ := primitive.ObjectIDFromHex(jobId)
	objCompany,_ := primitive.ObjectIDFromHex(companyId)
	result,err := db.Collection("companies").UpdateOne(ctx, bson.M{
		"$and":bson.A{
			bson.M{"_id":objCompany},
			bson.M{"jobs.id":objJob},
		},

	}, bson.M{
		"$set": bson.M{
			"jobs.$.title":      request.Title,
			"jobs.$.detail":     request.Detail,
			"jobs.$.min_salary": request.MinSalary,
			"jobs.$.max_salary": request.MaxSalary,
		},
	})
	if err != nil {
		return false,nil
	}
	return result.ModifiedCount > 0, err
}

func (m manipulationJob) TmpTakeDown(db *mongo.Database, ctx context.Context,current response.Job,companyId string) (bool, error) {
	objCompany,_ := primitive.ObjectIDFromHex(companyId)
	result,err := db.Collection("companies").UpdateOne(ctx,bson.M{
		"$and":bson.A{
			bson.M{"_id":objCompany},
			bson.M{"jobs.id":current.Id},
		},
	},bson.M{
		"$set":bson.M{
			"jobs.$.status":!current.Status,
		},
	})
	if err != nil{
		return false,err
	}
	return result.ModifiedCount > 0,nil
}
