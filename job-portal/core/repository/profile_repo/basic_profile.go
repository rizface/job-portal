package profile_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BasicProfile interface {
	GetProfile(ctx context.Context,db *mongo.Database,userId string) *mongo.SingleResult
	UpdateProfile(ctx context.Context,db *mongo.Database,userId string, data bson.M) (bool,error)
}