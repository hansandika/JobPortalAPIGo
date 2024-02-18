package job

import (
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type JobRepo struct {
	Job JobDataRepoItf
}

func NewMasterRepo(db *infra.DbConnection, logger *logrus.Logger) JobRepo {
	return JobRepo{
		Job: newJobDataRepo(db, logger),
	}
}
