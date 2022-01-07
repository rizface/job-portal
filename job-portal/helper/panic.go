package helper

import "job-portal/app/exception"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicException(e exception.Exception,condition bool) {
	if condition {
		panic(e)
	}
}