package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Applicant struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	NamDepan     string             `json:"nama_depan" bson:"nama_depan"`
	NamaBelakang string             `json:"nama_belakang" bson:"nama_belakang"`
	Password     string             `bson:"password"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Status       bool               `json:"status" bson:"status"`
}
