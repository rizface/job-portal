package test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type filter struct {
	Keyword bson.M `bson:"$text"`
}

func TestA(t *testing.T) {
	var f filter
	var b bson.M
	f.Keyword = bson.M{}
	a,_ :=bson.Marshal(f)
	bson.Unmarshal(a,&b)
	fmt.Println(b)
}
