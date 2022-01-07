package profile_service

import "job-portal/app/model"

type BasicProfile interface {
	GetDetail(userId string) model.Object
	UpdateDetail(userId string, user model.Object) string
}
