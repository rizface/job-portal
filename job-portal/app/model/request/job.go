package request

import (
	"go.mongodb.org/mongo-driver/bson"
	"job-portal/app/exception"
	"job-portal/app/model"
	"time"
)

type MinQualification struct {
	Requirements []string `json:"requirements" bson:"requirements"`
	Prefered     []string `json:"prefered" bson:"prefered"`
}

type Job struct {
	Id               string           `json:"id"  bson:"_id"`
	Title            string           `json:"title"  bson:"title" validate:"required"`
	Location         string           `json:"location" bson:"location"`
	MinSalary        int64            `json:"min_salary" bson:"min_salary"`
	MaxSalary        int64            `json:"max_salary" bson:"max_salary"`
	Type             string           `json:"type" bson:"type"`
	JobDescription   string           `json:"job_description" bson:"job_description"`
	MinQualification MinQualification `json:"min_qualification" bson:"min_qualification"`
	Status           bool             `bson:"status"`
	CreatedAt        time.Time        `bson:"created_at"`
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
	return bsonObject
}
