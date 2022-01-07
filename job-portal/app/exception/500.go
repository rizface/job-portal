package exception

import "net/http"

type InternalServerError struct {
	Err string
}

func (is InternalServerError) Code() int {
	return http.StatusInternalServerError
}

func (is InternalServerError) Error() string {
	return is.Err
}

