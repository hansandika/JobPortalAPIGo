package handler

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/handler/application"
	"github.com/hansandika/go-job-portal-api/handler/job"
	"github.com/hansandika/go-job-portal-api/handler/user"
	"github.com/hansandika/go-job-portal-api/usecase"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	User        user.UserHandler
	Job         job.JobHandler
	Application application.ApplicationHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, log *logrus.Logger) Handler {
	return Handler{
		User:        user.NewCoreHandler(uc, conf, log),
		Job:         job.NewCoreHandler(uc, conf, log),
		Application: application.NewCoreHandler(uc, conf, log),
	}
}
