package profile_controller

import "net/http"

type BasicProfile interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	UpdateDetail(w http.ResponseWriter, r *http.Request)
}
