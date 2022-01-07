package middleware

import (
	"context"
	"job-portal/app/exception"
	"job-portal/helper"
	"net/http"
	"strings"
)

func CompanyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},strings.Contains(token,"Bearer") == false)
		items := strings.Split(token," ")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},len(items) != 2)
		claim := helper.VerifyCompanyToken(items[1])
		request = request.WithContext(context.WithValue(request.Context(),"company-data",claim))
		next.ServeHTTP(writer,request)
	})
}
