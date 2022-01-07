package auth_service

import "job-portal/app/model/request"

type Auth interface {
	Login(request request.Auth, collection string) string
	Register(request request.Auth, collection string) string
	Aktivasi(userId string) bool
}
