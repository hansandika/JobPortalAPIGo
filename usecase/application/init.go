package application

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	"github.com/sirupsen/logrus"
)

type ApplicationUsecase struct {
	Application ApplicationDataUsecaseItf
}

func NewUsecase(repo repository.Repo, conf *general.SectionService, database *infra.DbConnection, logger *logrus.Logger) ApplicationUsecase {
	return ApplicationUsecase{
		Application: newApplicationDataUsecase(repo, conf, logger, database),
	}
}
