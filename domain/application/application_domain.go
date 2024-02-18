package application

import (
	"time"

	"github.com/hansandika/go-job-portal-api/domain/job"
	"github.com/hansandika/go-job-portal-api/domain/user"
	"gorm.io/gorm"
)

type Application struct {
	ApplicationId    uint `gorm:"primary_key;auto_increment:true"`
	JobId            uint
	UserId           uint
	Job              job.Job   `gorm:"references:JobId"`
	User             user.User `gorm:"references:UserId"`
	Status           int       `gorm:"default:1"`
	IsActive         bool      `gorm:"default:true"`
	CvLink           string    `gorm:"type:varchar(255)"`
	CoverLetterLink  string    `gorm:"type:varchar(255)"`
	YearOfExperience int       `gorm:"default:0"`
	UpdatedAt        time.Time
	CreatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
