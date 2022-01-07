package middleware

import (
	"context"
	exception2 "job-portal/app/exception"
	"job-portal/helper"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				exception,ok := err.(exception2.Exception)
				if ok {
					helper.JsonWriter(writer,exception.Code(),exception.Error(),nil)
				} else {
					helper.JsonWriter(writer,500,err.(error).Error(),nil)
				}
			}
		}()
		request = request.WithContext(context.WithValue(request.Context(),"data","fariz"))
		next.ServeHTTP(writer,request)
	})
}
