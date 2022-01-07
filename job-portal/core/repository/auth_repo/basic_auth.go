package auth_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"job-portal/app/model/request"
)

type Auth interface {
	Register(db *mongo.Database,ctx context.Context,request request.Auth,collection string) (interface{},error)
	Login(db *mongo.Database,ctx context.Context,request request.Auth,collection string) *mongo.SingleResult
	Aktivasi(db *mongo.Database,ctx context.Context,userId string,collection string) bool
}
