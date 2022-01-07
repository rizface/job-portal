package validation

import (
	"github.com/go-playground/validator/v10"
	exception2 "job-portal/app/exception"
	"job-portal/helper"
)

func Validation(err error) {
	for _,v := range err.(validator.ValidationErrors) {
		helper.PanicException(exception2.BadRequest{Err:v.Field() + " " + errMsg[v.Tag()]},true)
	}
}