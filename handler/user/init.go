package user

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/usecase"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	User UserDataHandler
}

func NewCoreHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: newUserHandler(usecase, conf, logger),
	}
}
