package application

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/usecase"
	"github.com/sirupsen/logrus"
)

type ApplicationHandler struct {
	Application ApplicationDataHandler
}

func NewCoreHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) ApplicationHandler {
	return ApplicationHandler{
		Application: newApplicationHandler(usecase, conf, logger),
	}
}
