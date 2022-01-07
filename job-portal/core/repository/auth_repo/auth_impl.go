package auth_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
	"time"
)

type auth struct{}

func NewAuth() Auth {
	return &auth{}
}

func (a *auth) Register(db *mongo.Database, ctx context.Context, request request.Auth,collection string) (interface{}, error) {
	result,err := db.Collection(collection).InsertOne(ctx,bson.M{
		"email":request.Email,
		"password": request.Password,
		"level": request.Level,
		"status": false,
		"createdAt": time.Now(),
	})
	return result.InsertedID.(primitive.ObjectID).Hex(),err
}

func (a *auth) Login(db *mongo.Database, ctx context.Context, request request.Auth,collection string) *mongo.SingleResult {
	cursor := db.Collection(collection).FindOne(ctx,bson.M{
		"email":bson.M{"$eq":request.Email},
	})
	return cursor
}

func (a *auth) Aktivasi(db *mongo.Database, ctx context.Context, userId string, collection string) bool {
	objId,_ := primitive.ObjectIDFromHex(userId)
	result,_ := db.Collection(collection).UpdateOne(ctx,bson.M{"_id":objId},bson.M{"$set":bson.M{"status":true}})
	return result.ModifiedCount > 0
}


