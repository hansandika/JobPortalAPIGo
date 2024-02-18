package usecase

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	"github.com/hansandika/go-job-portal-api/usecase/application"
	"github.com/hansandika/go-job-portal-api/usecase/job"
	"github.com/hansandika/go-job-portal-api/usecase/user"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	User        user.UserUsecase
	Job         job.JobUsecase
	Application application.ApplicationUsecase
}

func NewUsecase(repo repository.Repo, conf *general.SectionService, database *infra.DbConnection, logger *logrus.Logger) Usecase {
	return Usecase{
		User:        user.NewUsecase(repo, conf, database, logger),
		Job:         job.NewUsecase(repo, conf, database, logger),
		Application: application.NewUsecase(repo, conf, database, logger),
	}
}
