package request

import (
	"go.mongodb.org/mongo-driver/bson"
	"job-portal/app/exception"
	"job-portal/app/model"
	"job-portal/helper"
)

type Company struct {
	NamaPerusahaan   string `json:"nama_perusahaan" bson:"nama_perusahaan" validate:"required"`
	Alamat           string `json:"alamat" bson:"alamat" validate:"required"`
	DetailPerusahaan string `json:"detail_perusahaan" bson:"detail_perusahaan" validate:"required"`
	LinkWebsite      string `json:"link_website" bson:"link_website"`
	_                struct{}
}

func NewCompany() model.Object {
	return &Company{}
}

func (c *Company) Convert() bson.M {
	var final bson.M
	bytesResult, err := bson.Marshal(c)
	helper.PanicException(exception.InternalServerError{Err: "terjadi kesalahan pada sistem kami"}, err != nil)
	err = bson.Unmarshal(bytesResult, &final)
	return final
}
