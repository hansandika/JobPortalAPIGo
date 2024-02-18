package user

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	User UserDataUsecaseItf
}

func NewUsecase(repo repository.Repo, conf *general.SectionService, database *infra.DbConnection, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		User: newUserDataUsecase(repo, conf, logger, database),
	}
}
