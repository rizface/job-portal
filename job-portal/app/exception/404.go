package exception

import "net/http"

type NotFouund struct {
	Err string
}

func (nf NotFouund) Code() int {
	return http.StatusNotFound
}

func (nf NotFouund) Error() string {
	return nf.Err
}

