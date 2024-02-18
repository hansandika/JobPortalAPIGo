package job

import (
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/usecase"
	"github.com/sirupsen/logrus"
)

type JobHandler struct {
	Job JobDataHandler
}

func NewCoreHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) JobHandler {
	return JobHandler{
		Job: newJobHandler(usecase, conf, logger),
	}
}
