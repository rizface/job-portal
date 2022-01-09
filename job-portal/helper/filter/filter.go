package filter_builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"job-portal/app/exception"
	"job-portal/app/model/request"
	"job-portal/helper"
	"net/http"
	"os"
	"strconv"
)

func buildPaging(filter *request.Filter, start string) {
	filter.Limit = 15
	if len(start) == 0 || start == "0" {
		filter.Skip = 0
	} else {
		v,err := strconv.Atoi(start)
		helper.PanicException(exception.NotFound{Err: "halaman yang kamu cari tidak ditemukan"}, err != nil)
		filter.Skip = v - 1
	}
}

func GetPrevNext(length,skip,limit int) (string,string) {
	if length != limit {
		return "prev",os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")+"/jobs"
	} else {
		if skip > 0 {
			skip += 1
		}
		return "next",os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")+"/jobs?start="+strconv.Itoa(skip + 15)
	}
}

func buildSearch(filter *request.Filter, keyword string){
	if len(keyword) == 0 {
		filter.Keyword = bson.M{}
	} else {
		filter.Keyword = bson.M{"$search":keyword}
	}
}

func BuildFilter(r *http.Request) *request.Filter {
	filter := new(request.Filter)
	buildPaging(filter, helper.GetQuery(r,"start"))
	buildSearch(filter,helper.GetQuery(r,"search"))
	return filter
}
