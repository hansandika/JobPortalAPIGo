package application

import (
	"net/http"

	dto "github.com/hansandika/go-job-portal-api/domain/application"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/exception"
	utils "github.com/hansandika/go-job-portal-api/global/utils/general"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	r "github.com/hansandika/go-job-portal-api/repository/application"
	rj "github.com/hansandika/go-job-portal-api/repository/job"
	ru "github.com/hansandika/go-job-portal-api/repository/user"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ApplicationDataUsecaseItf interface {
	CreateApplication(req *dto.ApplicationAddRequest) (*dto.ApplicationResponse, *general.Error)
	UpdateApplication(req *dto.ApplicationUpdateRequest) (*dto.ApplicationResponse, *general.Error)
	GetApplicationByUserEmailAndJobId(req *dto.GetApplicationRequest) ([]dto.ApplicationResponse, *general.Error)
	GetJobApplicationsByJobId(jobId uint) ([]dto.ApplicationResponse, *general.Error)
}

type ApplicationDataUsecase struct {
	Repo     r.ApplicationDataRepoItf
	RepoUser ru.UserDataRepoItf
	RepoJob  rj.JobDataRepoItf
	Database *infra.DbConnection
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newApplicationDataUsecase(r repository.Repo, conf *general.SectionService, logger *logrus.Logger, database *infra.DbConnection) ApplicationDataUsecase {
	return ApplicationDataUsecase{
		Repo:     r.Application.Application,
		RepoUser: r.User.User,
		RepoJob:  r.Job.Job,
		Conf:     conf,
		Log:      logger,
		Database: database,
	}
}

func (a ApplicationDataUsecase) CreateApplication(req *dto.ApplicationAddRequest) (*dto.ApplicationResponse, *general.Error) {
	// Validate request
	if err := req.Validate(); err != nil {
		a.Log.Error(err)
		return nil, err
	}

	// Get User Detail From Email Inside Cookie
	userEmail, err := a.RepoUser.GetUserByEmail(req.Email)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get user by email")
	}

	if userEmail == nil {
		return nil, exception.NewError(http.StatusNotFound, "User not found")
	}

	// Check if job exists
	job, err := a.RepoJob.GetJobByJobId(req.JobId)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get job by id")
	}

	if job == nil {
		return nil, exception.NewError(http.StatusNotFound, "Job not found")
	}

	// Check if user has already applied for the job
	application, err := a.Repo.GetActiveApplicationByUserIdAndJobId(userEmail.UserId, job.JobId)

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get application by user id and job id")
	}

	if application != nil {
		return nil, exception.NewError(http.StatusBadRequest, "User has already applied for the job")
	}

	// Create application object
	application = &dto.Application{
		JobId:            job.JobId,
		UserId:           userEmail.UserId,
		Status:           req.Status,
		IsActive:         utils.GenerateBool(req.IsActive),
		CvLink:           req.CvLink,
		CoverLetterLink:  req.CoverLetterLink,
		YearOfExperience: req.YearOfExperience,
	}

	// Insert application
	application, err = a.Repo.Insert(application)
	if err != nil {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to insert application")
	}

	// Create response
	applicationResponse := &dto.ApplicationResponse{
		ApplicationId:    application.ApplicationId,
		JobId:            application.JobId,
		JobTitle:         job.JobTitle,
		UserId:           application.UserId,
		Email:            userEmail.Email,
		Name:             userEmail.Name,
		Status:           application.Status,
		StatusName:       dto.GetApplicationStatusName(application.Status),
		IsActive:         utils.GenerateInt(application.IsActive),
		CvLink:           application.CvLink,
		CoverLetterLink:  application.CoverLetterLink,
		YearOfExperience: application.YearOfExperience,
		IssuedAt:         application.CreatedAt,
	}

	return applicationResponse, nil
}

func (a ApplicationDataUsecase) UpdateApplication(req *dto.ApplicationUpdateRequest) (*dto.ApplicationResponse, *general.Error) {
	// Set IsActive based on Status
	req.IsActive = dto.GetApplicationActiveBaseOnStatus(req.Status)

	// Validate request
	if err := req.Validate(); err != nil {
		a.Log.Error(err)
		return nil, err
	}

	// Get User Detail From Email Inside Cookie
	userEmail, err := a.RepoUser.GetUserByEmail(req.Email)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get user by email")
	}

	if userEmail == nil {
		return nil, exception.NewError(http.StatusNotFound, "User not found")
	}

	// Check if job exists
	job, err := a.RepoJob.GetJobByJobId(req.JobId)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get job by id")
	}

	if job == nil {
		return nil, exception.NewError(http.StatusNotFound, "Job not found")
	}

	// Check if user has already applied for the job
	application, err := a.Repo.GetActiveApplicationByUserIdAndJobId(userEmail.UserId, job.JobId)

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get application by user id and job id")
	}

	if application == nil {
		return nil, exception.NewError(http.StatusBadRequest, "User has not applied for the job")
	}

	// Update application status
	application.Status = req.Status
	application.IsActive = utils.GenerateBool(req.IsActive)
	application, err = a.Repo.UpdateStatus(application)
	if err != nil {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to update application status")
	}

	// Create response
	applicationResponse := &dto.ApplicationResponse{
		ApplicationId:    application.ApplicationId,
		JobId:            application.JobId,
		JobTitle:         job.JobTitle,
		UserId:           application.UserId,
		Email:            userEmail.Email,
		Name:             userEmail.Name,
		Status:           application.Status,
		StatusName:       dto.GetApplicationStatusName(application.Status),
		IsActive:         utils.GenerateInt(application.IsActive),
		CvLink:           application.CvLink,
		CoverLetterLink:  application.CoverLetterLink,
		YearOfExperience: application.YearOfExperience,
		IssuedAt:         application.CreatedAt,
	}

	return applicationResponse, nil
}

func (a ApplicationDataUsecase) GetApplicationByUserEmailAndJobId(req *dto.GetApplicationRequest) ([]dto.ApplicationResponse, *general.Error) {
	// Validate request
	if err := req.Validate(); err != nil {
		a.Log.Error(err)
		return nil, err
	}

	// Get User Detail
	userEmail, err := a.RepoUser.GetUserByEmail(req.Email)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get user by email")
	}

	if userEmail == nil {
		return nil, exception.NewError(http.StatusNotFound, "User not found")
	}

	// Check if job exists
	job, err := a.RepoJob.GetJobByJobId(req.JobId)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get job by id")
	}

	if job == nil {
		return nil, exception.NewError(http.StatusNotFound, "Job not found")
	}

	// Get application by user id and job id
	applications, err := a.Repo.GetAllApplicationByUserIdAndJobId(userEmail.UserId, req.JobId)

	if err != nil {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get applications by user id and job id")
	}

	// Check if applications is empty
	if len(applications) == 0 {
		return nil, exception.NewError(http.StatusNotFound, "Application not found")
	}

	// Create response
	applicationResponses := make([]dto.ApplicationResponse, 0)

	for _, application := range applications {
		applicationResponse := dto.ApplicationResponse{
			ApplicationId:    application.ApplicationId,
			JobId:            application.JobId,
			JobTitle:         job.JobTitle,
			UserId:           application.UserId,
			Email:            userEmail.Email,
			Name:             userEmail.Name,
			Status:           application.Status,
			StatusName:       dto.GetApplicationStatusName(application.Status),
			IsActive:         utils.GenerateInt(application.IsActive),
			CvLink:           application.CvLink,
			CoverLetterLink:  application.CoverLetterLink,
			YearOfExperience: application.YearOfExperience,
			IssuedAt:         application.CreatedAt,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	return applicationResponses, nil
}

func (a ApplicationDataUsecase) GetJobApplicationsByJobId(jobId uint) ([]dto.ApplicationResponse, *general.Error) {
	// Validate request
	if jobId == 0 {
		return nil, exception.NewError(http.StatusBadRequest, "JobId is required")
	}

	// Check if job exists
	job, err := a.RepoJob.GetJobByJobId(jobId)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get job by id")
	}

	// Get all applications by job id
	applications, err := a.Repo.GetAllJobApplicationByJobId(jobId)
	if err != nil {
		a.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get applications by job id")
	}

	// Check if applications is empty
	if len(applications) == 0 {
		return nil, exception.NewError(http.StatusNotFound, "Application not found")
	}

	// Create response
	applicationResponses := make([]dto.ApplicationResponse, 0)
	for _, application := range applications {
		applicationResponse := dto.ApplicationResponse{
			ApplicationId:    application.ApplicationId,
			JobId:            application.JobId,
			JobTitle:         job.JobTitle,
			UserId:           application.UserId,
			Email:            application.User.Email,
			Name:             application.User.Name,
			Status:           application.Status,
			StatusName:       dto.GetApplicationStatusName(application.Status),
			IsActive:         utils.GenerateInt(application.IsActive),
			CvLink:           application.CvLink,
			CoverLetterLink:  application.CoverLetterLink,
			YearOfExperience: application.YearOfExperience,
			IssuedAt:         application.CreatedAt,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	return applicationResponses, nil
}
