package auth_controller

import "net/http"

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Aktivasi(w http.ResponseWriter, r *http.Request)
}
