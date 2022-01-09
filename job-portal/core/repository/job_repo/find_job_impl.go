package job_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"job-portal/app/model/request"
	"job-portal/core/repository/job_repo/job_interface"
)

type findJob struct {}

func NewFindJob() job_interface.ApplicantFindJob {
	return &findJob{}
}

func (f findJob) GetJob(db *mongo.Database, ctx context.Context,filter request.Filter) (*mongo.Cursor, error) {
	cursor,err := db.Collection("jobs").Find(ctx,filter.Convert(),options.Find().SetProjection(bson.M{
		"title":1,
		"company":1,
		"location":1,
		"min_salary":1,
		"max_salary":1,
		"type":1,
		"status":1,
		"created_at":1,
		"score":bson.M{"$meta":"textScore"},
	}).SetSkip(int64(filter.Skip)).SetLimit(int64(filter.Limit)).SetSort(bson.M{"score":-1}))
	return cursor,err
}

func (f findJob) DetailJob(db *mongo.Database, ctx context.Context, jobId string) *mongo.SingleResult {
	objJob,_ := primitive.ObjectIDFromHex(jobId)
	cursor := db.Collection("jobs").FindOne(ctx,bson.M{
		"_id": objJob,
	},options.FindOne().SetProjection(bson.M{
		"applicants":0,
	}))
	return cursor
}

func (f findJob) JobRecomendation(db *mongo.Database, ctx context.Context, interset []string) (*mongo.Cursor, error) {
	panic("panic me")
}



