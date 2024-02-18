package user

import (
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	User UserDataRepoItf
}

func NewMasterRepo(db *infra.DbConnection, logger *logrus.Logger) UserRepo {
	return UserRepo{
		User: newUserDataRepo(db, logger),
	}
}
