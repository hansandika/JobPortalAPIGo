package application

import (
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type ApplicationRepo struct {
	Application ApplicationDataRepoItf
}

func NewMasterRepo(db *infra.DbConnection, logger *logrus.Logger) ApplicationRepo {
	return ApplicationRepo{
		Application: newApplicationDataRepo(db, logger),
	}
}
