package middleware

import (
	"job-portal/app/exception"
	"job-portal/helper"
	"net/http"
	"strings"
)

func CompanyCompleteProfile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},strings.Contains(token,"Bearer") == false)
		items := strings.Split(token," ")
		helper.PanicException(exception.UnAuthorized{Err:"token kamu tidak valid"},len(items) != 2)
		claim := helper.VerifyCompanyToken(items[1])
		data := claim.(*helper.Company)
		if data.NamaPerusahaan == "" {
			helper.PanicException(exception.Forbidden{Err:"kamu tidak punya akses ke fitur ini, lengkapi profile terlebih dahulu"}, true)
			return
		}
		next.ServeHTTP(writer,request)
	})
}
