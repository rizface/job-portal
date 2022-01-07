package model

import "go.mongodb.org/mongo-driver/bson"

type Object interface {
	Convert() bson.M
}
