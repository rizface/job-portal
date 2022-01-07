package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	NamaPerusahaan string             `json:"nama_perusahaan" bson:"nama_perusahaan"`
	Password       string             `bson:"password"`
	Email          string             `json:"email" bson:"email"`
	Username       string             `json:"username" bson:"username"`
	Status         bool               `json:"status" bson:"status"`
}
