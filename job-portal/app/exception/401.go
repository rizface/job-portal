package exception

import "net/http"

type UnAuthorized struct {
	Err string
}

func (un UnAuthorized) Code() int {
	return http.StatusUnauthorized
}

func (un UnAuthorized) Error() string {
	return un.Err
}

