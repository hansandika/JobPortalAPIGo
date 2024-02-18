package infra

import (
	"fmt"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Database *gorm.DB
}

func NewConnection(conf *general.SectionService) *DbConnection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Database.DBHost, conf.Database.Username, conf.Database.Password, conf.Database.DBName, conf.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
		return nil
	}

	return &DbConnection{db}
}
