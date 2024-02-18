package job

import (
	dj "github.com/hansandika/go-job-portal-api/domain/job"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type JobDataRepo struct {
	DB  *infra.DbConnection
	Log *logrus.Logger
}

func newJobDataRepo(db *infra.DbConnection, logger *logrus.Logger) JobDataRepo {
	return JobDataRepo{
		DB:  db,
		Log: logger,
	}
}

type JobDataRepoItf interface {
	GetJobByJobId(jobId uint) (*dj.Job, error)
	GetJobByJobTitle(jobTitle string) (*dj.Job, error)
	Insert(job *dj.Job) (*dj.Job, error)
	ListJobs() ([]dj.Job, error)
}

func (jr JobDataRepo) GetJobByJobId(jobId uint) (*dj.Job, error) {
	var job dj.Job
	if err := jr.DB.Database.Table("jobs").Where("job_id = ?", jobId).First(&job).Error; err != nil {
		jr.Log.Error(err)
		return nil, err
	}

	return &job, nil
}

func (jr JobDataRepo) GetJobByJobTitle(jobTitle string) (*dj.Job, error) {
	var job dj.Job
	if err := jr.DB.Database.Table("jobs").Where("job_title = ?", jobTitle).First(&job).Error; err != nil {
		jr.Log.Error(err)
		return nil, err
	}

	return &job, nil
}

func (jr JobDataRepo) Insert(job *dj.Job) (*dj.Job, error) {
	if err := jr.DB.Database.Create(&job).Error; err != nil {
		jr.Log.Error(err)
		return nil, err
	}

	return job, nil
}

func (jr JobDataRepo) ListJobs() ([]dj.Job, error) {
	var jobs []dj.Job
	if err := jr.DB.Database.Find(&jobs).Error; err != nil {
		jr.Log.Error(err)
		return nil, err
	}

	return jobs, nil
}
