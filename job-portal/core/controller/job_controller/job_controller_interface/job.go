package job_controller_interface

import "net/http"

type Manipulation interface {
	PostJob(w http.ResponseWriter, r *http.Request)
	DetailJob(w http.ResponseWriter, r *http.Request)
	DeleteJob(w http.ResponseWriter, r *http.Request)
	UpdateJob(w http.ResponseWriter, r *http.Request)
	TakeDown(w http.ResponseWriter, r *http.Request)
}