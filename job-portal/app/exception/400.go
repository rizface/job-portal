package exception

import "net/http"

type BadRequest struct {
	Err string
}

func (b BadRequest) Code() int {
	return http.StatusBadRequest
}

func (b BadRequest) Error() string {
	return b.Err
}

