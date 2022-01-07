package request

import (
	"go.mongodb.org/mongo-driver/bson"
	"job-portal/app/exception"
	"job-portal/app/model"
	"job-portal/helper"
)

type Applicant struct {
	NamaDepan          string `json:"nama_depan" bson:"nama_depan" validate:"required"`
	NamaBelakang       string `json:"nama_belakang" bson:"nama_belakang" validate:"required"`
	PendidikanTerakhir string `json:"pendidikan_terakhir" bson:"pendidikan_terakhir" validate:"required"`
	TentangSaya        string `json:"tentang_saya" bson:"tentang_saya" validate:"required"`
	_                  struct{}
}

func NewApplicant() model.Object {
	return &Applicant{}
}

func (a *Applicant) Convert() bson.M {
	var final bson.M
	bytesResult, err := bson.Marshal(a)
	helper.PanicException(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"}, err != nil)
	err = bson.Unmarshal(bytesResult, &final)
	return final
}
