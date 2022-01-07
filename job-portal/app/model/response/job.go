package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Job struct {
	Id        primitive.ObjectID `json:"id" bson:"id"`
	Title     string             `json:"title"  bson:"title""`
	Detail    string             `json:"detail" bson:"detail"`
	MinSalary interface{}        `json:"min_salary" bson:"min_salary"`
	MaxSalary interface{}        `json:"max_salary" bson:"max_salary"`
	Status    bool               `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
