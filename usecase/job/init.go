package job

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	"github.com/sirupsen/logrus"
)

type JobUsecase struct {
	Job JobDataUsecaseItf
}

func NewUsecase(repo repository.Repo, conf *general.SectionService, database *infra.DbConnection, logger *logrus.Logger) JobUsecase {
	return JobUsecase{
		Job: newJobDataUsecase(repo, conf, logger, database),
	}
}
