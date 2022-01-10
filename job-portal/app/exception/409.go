package exception

import "net/http"

type Conflict struct {
	Err string
}

func (c Conflict) Code() int {
	return http.StatusConflict
}

func (c Conflict) Error() string {
	return c.Err
}

