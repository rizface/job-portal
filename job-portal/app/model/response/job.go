package response

import (
	"time"
)

type MinQualification struct {
	Requirements []string `json:"requirements" bson:"requirements"`
	Prefered     []string `json:"prefered" bson:"prefered"`
}

type Job struct {
	Id               string           `json:"id,omitempty"  bson:"_id"`
	Title            string           `json:"title,omitempty"  bson:"title" validate:"required"`
	Location         string           `json:"location,omitempty" bson:"location"`
	MinSalary        int64            `json:"min_salary,omitempty" bson:"min_salary"`
	MaxSalary        int64            `json:"max_salary,omitempty" bson:"max_salary"`
	Type             string           `json:"type,omitempty" bson:"type"`
	JobDescription   string           `json:"job_description,omitempty" bson:"job_description"`
	MinQualification *MinQualification `json:"min_qualification,omitempty" bson:"min_qualification"`
	Applicants       []Applicant      `json:"applicants,omitempty" bson:"applicants"`
	Status           bool             `bson:"status"`
	CreatedAt        time.Time        `bson:"created_at"`
}
