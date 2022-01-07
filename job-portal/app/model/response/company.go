package response

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"job-portal/app/exception"
	"job-portal/app/model"
)

type Company struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Email            string             `json:"email" bson:"email"`
	NamaPerusahaan   string             `json:"nama_perusahaan" bson:"nama_perusahaan"`
	DetailPerusahaan string             `json:"detail_perusahaan" bson:"detail_perusahaan"`
	LinkWebsite      string             `json:"link_website" bson:"link_website"`
	Alamat           string             `json:"alamat" bson:"alamat"`
	Jobs             []Job              `json:"jobs,omitempty" bson:"jobs,omitempty"`
	Password         string             `json:"password,omitempty" bson:"password,omitempty"`
	Status           bool               `json:"status" bson:"status"`
}

func NewCompany() model.Object {
	return &Company{}
}

func (c *Company) Convert() bson.M {
	var final bson.M
	bytesResult, err := bson.Marshal(c)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	err = bson.Unmarshal(bytesResult, &final)
	if err != nil {
		panic(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"})
	}
	return final
}
