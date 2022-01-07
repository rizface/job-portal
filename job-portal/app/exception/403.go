package exception

import "net/http"

type Forbidden struct {
	Err string
}

func (f Forbidden) Code() int {
	return http.StatusForbidden
}

func (f Forbidden) Error() string {
	return f.Err
}

