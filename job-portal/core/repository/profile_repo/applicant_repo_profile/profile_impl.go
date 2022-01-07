package applicant_repo_profile

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"job-portal/app/exception"
	"job-portal/core/repository/profile_repo"
	"job-portal/helper"
)

type profile struct{}

func NewProfile() profile_repo.BasicProfile {
	return &profile{}
}

func (p profile) GetProfile(ctx context.Context, db *mongo.Database, userId string) *mongo.SingleResult {
	objId,err := primitive.ObjectIDFromHex(userId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	cursor := db.Collection("applicants").FindOne(ctx,bson.M{"_id":bson.M{"$eq":objId}},options.FindOne().SetProjection(bson.M{
		"password":0,
	}))
	return cursor
}

func (p profile) UpdateProfile(ctx context.Context, db *mongo.Database, userId string, data bson.M) (bool, error) {
	objId,err := primitive.ObjectIDFromHex(userId)
	helper.PanicException(exception.InternalServerError{Err:"terjadi kesalahan pada sistem kami"}, err != nil)
	result,err := db.Collection("applicants").UpdateOne(ctx,bson.M{
		"_id":bson.M{"$eq":objId},
	},bson.M{
		"$set":bson.M{
		"nama_depan": data["nama_depan"],
		"nama_belakang": data["nama_belakang"],
		"pendidikan_terakhir": data["pendidikan_terakhir"],
		"tentang_saya": data["tentang_saya"],
	}})
	if err != nil {
		return false,err
	}
	return result.ModifiedCount > 0,nil
}


