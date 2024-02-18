package job

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Job struct {
	JobId       uint `gorm:"primaryKey"`
	JobTitle    string
	Description string         `gorm:"unique"`
	Requirement pq.StringArray `gorm:"type:text[]"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
