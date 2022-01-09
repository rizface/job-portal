package request

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Skip    int    `bson:"skip,omitempty"`
	Limit   int    `bson:"limit,omitempty"`
	Keyword bson.M `bson:"$text,omitempty"`
}

func (f *Filter) Convert() bson.M {
	newFilter := Filter{
		Keyword: f.Keyword,
	}
	var result bson.M
	bsonBytes, _ := bson.Marshal(newFilter)
	bson.Unmarshal(bsonBytes, &result)
	return result
}
