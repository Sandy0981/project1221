package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
	UserId      string `json:"user_id"`
	Jobs        []Job  `json:"jobs" gorm:"foreignKey:CompanyID"`
}

type NewCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	FoundedYear string `json:"founded_year" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Jobs        []Job  `json:"jobs"`
}

type Job struct {
	gorm.Model
	Title           string `json:"title"`
	ExperienceLevel string `json:"experience_required"`
	CompanyID       uint   `json:"company_id"`
}

/*
"company_name":"TekSystem",
"founded_year":"2010",
"location":"bangalore",
"user_id":"121",
"jobs": [
			{
				"title": "Software Engineer",
				"department": "Engineering",
				"company_id": 1
			},
			{
				"title": "UX Designer",
				"department": "Design",
				"company_id": 1
			}
		]
	}`


*/
