package job

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	dto "github.com/hansandika/go-job-portal-api/domain/job"
	utils "github.com/hansandika/go-job-portal-api/global/utils/general"
	uj "github.com/hansandika/go-job-portal-api/usecase/job"

	"github.com/hansandika/go-job-portal-api/usecase"
	"github.com/sirupsen/logrus"
)

type JobDataHandler struct {
	Usecase uj.JobDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newJobHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) JobDataHandler {
	return JobDataHandler{
		Usecase: usecase.Job.Job,
		conf:    conf,
		log:     logger,
	}
}

func (jh JobDataHandler) CreateJob(c *gin.Context) {
	var payload dto.JobAddRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Error Parsing Json", nil, nil)
		return
	}

	if err := utils.JobAddValidator.Validate(payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Request", nil, err)
		return
	}

	// Usecase
	jobResponse, err := jh.Usecase.CreateJob(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Create Job", jobResponse, nil)
}

func (jh JobDataHandler) ListJobs(c *gin.Context) {
	// Usecase
	jobResponse, err := jh.Usecase.ListJobs()
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To List Jobs", jobResponse, nil)
}
