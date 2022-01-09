package middleware

import (
	"context"
	"fmt"
	"job-portal/app/exception"
	"job-portal/helper"
	"net/http"
	"strings"
)

func MixAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},strings.Contains(token,"Bearer") == false)
		items := strings.Split(token," ")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},len(items) != 2)
		claim := helper.VerifyApplicantToken(items[1])
		level := claim.(*helper.Applicant).Level
		fmt.Println(level)
		if  level != "applicant" && level != "company"{
			panic(exception.Forbidden{Err:"kamu tidak punya akses kesini"})
		}
		request = request.WithContext(context.WithValue(request.Context(),"mix-data",claim))
		next.ServeHTTP(writer,request)
	})
}
