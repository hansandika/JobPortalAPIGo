package repository

import (
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository/application"
	"github.com/hansandika/go-job-portal-api/repository/job"
	"github.com/hansandika/go-job-portal-api/repository/user"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	User        user.UserRepo
	Job         job.JobRepo
	Application application.ApplicationRepo
}

func NewRepo(database *infra.DbConnection, logger *logrus.Logger) Repo {
	return Repo{
		User:        user.NewMasterRepo(database, logger),
		Job:         job.NewMasterRepo(database, logger),
		Application: application.NewMasterRepo(database, logger),
	}
}
