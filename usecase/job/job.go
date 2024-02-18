package job

import (
	"net/http"

	"github.com/hansandika/go-job-portal-api/domain/general"
	dto "github.com/hansandika/go-job-portal-api/domain/job"
	"github.com/hansandika/go-job-portal-api/global/exception"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	r "github.com/hansandika/go-job-portal-api/repository/job"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JobDataUsecaseItf interface {
	CreateJob(req *dto.JobAddRequest) (*dto.JobResponse, *general.Error)
	ListJobs() ([]dto.JobResponse, *general.Error)
}

type JobDataUsecase struct {
	Repo     r.JobDataRepoItf
	Database *infra.DbConnection
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newJobDataUsecase(r repository.Repo, conf *general.SectionService, logger *logrus.Logger, database *infra.DbConnection) JobDataUsecase {
	return JobDataUsecase{
		Repo:     r.Job.Job,
		Conf:     conf,
		Log:      logger,
		Database: database,
	}
}

func (ut JobDataUsecase) CreateJob(req *dto.JobAddRequest) (*dto.JobResponse, *general.Error) {
	// Validate request
	if err := req.Validate(); err != nil {
		ut.Log.Error(err)
		return nil, err
	}

	// Check if job title already exist
	job, err := ut.Repo.GetJobByJobTitle(req.JobTitle)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get job by job title")
	}

	if job != nil {
		return nil, exception.NewError(http.StatusBadRequest, "Job title already exist")
	}

	// Create job object
	job = &dto.Job{
		JobTitle:    req.JobTitle,
		Description: req.Description,
		Requirement: req.Requirement,
	}

	// Insert job
	job, err = ut.Repo.Insert(job)
	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to create job")
	}

	// Create response
	jobResponse := &dto.JobResponse{
		JobID:       job.JobId,
		JobTitle:    job.JobTitle,
		Description: job.Description,
		Requirement: job.Requirement,
	}

	return jobResponse, nil
}

func (ut JobDataUsecase) ListJobs() ([]dto.JobResponse, *general.Error) {
	// Get list of jobs
	jobs, err := ut.Repo.ListJobs()
	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to get list of jobs")
	}

	// Create response
	var jobResponse []dto.JobResponse
	for _, job := range jobs {
		jobResponse = append(jobResponse, dto.JobResponse{
			JobID:       job.JobId,
			JobTitle:    job.JobTitle,
			Description: job.Description,
			Requirement: job.Requirement,
		})
	}

	return jobResponse, nil
}
