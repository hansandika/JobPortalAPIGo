package application

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/hansandika/go-job-portal-api/domain/application"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	utils "github.com/hansandika/go-job-portal-api/global/utils/general"
	"github.com/hansandika/go-job-portal-api/usecase"
	ua "github.com/hansandika/go-job-portal-api/usecase/application"
	"github.com/sirupsen/logrus"
)

type ApplicationDataHandler struct {
	Usecase ua.ApplicationDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newApplicationHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) ApplicationDataHandler {
	return ApplicationDataHandler{
		Usecase: usecase.Application.Application,
		conf:    conf,
		log:     logger,
	}
}

func (ah ApplicationDataHandler) CreateApplication(c *gin.Context) {
	var payload dto.ApplicationAddRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Error Parsing Json", nil, nil)
		return
	}

	// Get Job Id from Param
	jobId := c.Param("jobId")

	// Validate Job Id
	jobIdInt, errReq := strconv.Atoi(jobId)
	if errReq != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Job Id Param", nil, nil)
		return
	}

	// Get User Email from Cookie
	email, cookieExist := c.Get("email")
	if !cookieExist {
		general.CreateResponse(c, http.StatusBadRequest, "Error Getting Email From Cookie", nil, nil)
		return
	}

	// Set Payload
	payload.Email = email.(string)
	payload.Status = constant.OnHoldStatusId
	payload.JobId = uint(jobIdInt)
	payload.IsActive = constant.UserStatusActiveId

	if err := utils.ApplicationAddValidator.Validate(payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Request", nil, err)
		return
	}

	// Usecase
	applicationResponse, err := ah.Usecase.CreateApplication(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Create Application", applicationResponse, nil)
}

func (ah ApplicationDataHandler) UpdateApplication(c *gin.Context) {
	var payload dto.ApplicationUpdateRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Error Parsing Json", nil, nil)
		return
	}

	if err := utils.ApplicationUpdateValidator.Validate(payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Request", nil, err)
		return
	}

	// Usecase
	applicationResponse, err := ah.Usecase.UpdateApplication(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Update Application", applicationResponse, nil)
}

func (ah ApplicationDataHandler) GetLoggedUserApplication(c *gin.Context) {
	// Get Job Id from Param
	jobId := c.Param("jobId")

	// Validate Job Id
	jobIdInt, errReq := strconv.Atoi(jobId)
	if errReq != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Job Id Param", nil, nil)
		return
	}

	// Get User Email from Cookie
	email, cookieExist := c.Get("email")
	if !cookieExist {
		general.CreateResponse(c, http.StatusBadRequest, "Error Getting Email From Cookie", nil, nil)
		return
	}

	// Request
	var payload dto.GetApplicationRequest
	payload.Email = email.(string)
	payload.JobId = uint(jobIdInt)

	// Usecase
	applicationResponse, err := ah.Usecase.GetApplicationByUserEmailAndJobId(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Get Application By User", applicationResponse, nil)
}

func (ah ApplicationDataHandler) GetJobApplicationsByUserEmailAndJobId(c *gin.Context) {
	// Get Job Id from Param
	jobId := c.Param("jobId")

	// Validate Job Id
	jobIdInt, errReq := strconv.Atoi(jobId)
	if errReq != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Job Id Param", nil, nil)
		return
	}

	// Get User Email from Param
	email := c.Param("email")

	// Request
	var payload dto.GetApplicationRequest
	payload.Email = email
	payload.JobId = uint(jobIdInt)

	// Usecase
	applicationResponse, err := ah.Usecase.GetApplicationByUserEmailAndJobId(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Get Application By User", applicationResponse, nil)
}

func (ah ApplicationDataHandler) GetJobApplicationsByJobId(c *gin.Context) {
	// Get Job Id from Param
	jobId := c.Param("jobId")

	// Validate Job Id
	jobIdInt, errReq := strconv.Atoi(jobId)
	if errReq != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Job Id Param", nil, nil)
		return
	}

	// Usecase
	applicationResponse, err := ah.Usecase.GetJobApplicationsByJobId(uint(jobIdInt))
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Get Job Applications", applicationResponse, nil)
}
