package exception

import "net/http"

type NotFound struct {
	Err string
}

func (nf NotFound) Code() int {
	return http.StatusNotFound
}

func (nf NotFound) Error() string {
	return nf.Err
}

