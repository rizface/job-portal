package response

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/exception"
	"job-portal/app/model"
)

type Applicant struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	NamDepan     string             `json:"nama_depan" bson:"nama_depan"`
	NamaBelakang string             `json:"nama_belakang" bson:"nama_belakang"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Email        string             `json:"email" bson:"email"`
	Level        string             `json:"level" bson:"level"`
	Status       bool               `json:"status" bson:"status"`
}

func NewApplicant() model.Object {
	return &Applicant{}
}

func (a *Applicant) Convert() bson.M {
	var final bson.M
	bytesResult, err := bson.Marshal(a)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	err = bson.Unmarshal(bytesResult, &final)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	return final
}
