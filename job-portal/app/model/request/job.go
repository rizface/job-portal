package request

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/exception"
	"job-portal/app/model"
	"time"
)

type Job struct {
	Title     string    `json:"title"  bson:"title" validate:"required"`
	Detail    string    `json:"detail" bson:"detail" validate:"required"`
	MinSalary int64     `json:"min_salary" bson:"min_salary"`
	MaxSalary int64     `json:"max_salary" bson:"max_salary"`
	Status    bool      `bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewJob() model.Object {
	return &Job{}
}

func (j *Job) Convert() bson.M {
	var bsonObject bson.M
	j.CreatedAt = time.Now()
	j.Status = true
	bsonBytes, err := bson.Marshal(j)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	err = bson.Unmarshal(bsonBytes, &bsonObject)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	bsonObject["id"] = primitive.NewObjectID()
	return bsonObject
}
