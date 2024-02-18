package application

import (
	dp "github.com/hansandika/go-job-portal-api/domain/application"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type ApplicationDataRepo struct {
	DB  *infra.DbConnection
	Log *logrus.Logger
}

func newApplicationDataRepo(db *infra.DbConnection, logger *logrus.Logger) ApplicationDataRepo {
	return ApplicationDataRepo{
		DB:  db,
		Log: logger,
	}
}

type ApplicationDataRepoItf interface {
	Insert(Application *dp.Application) (*dp.Application, error)
	GetActiveApplicationByUserIdAndJobId(UserId uint, JobId uint) (*dp.Application, error)
	GetAllApplicationByUserIdAndJobId(UserId uint, JobId uint) ([]dp.Application, error)
	GetAllJobApplicationByJobId(JobId uint) ([]dp.Application, error)
	UpdateStatus(Application *dp.Application) (*dp.Application, error)
}

func (ar ApplicationDataRepo) Insert(application *dp.Application) (*dp.Application, error) {
	if err := ar.DB.Database.Create(&application).Error; err != nil {
		ar.Log.Error(err)
		return nil, err
	}

	return application, nil
}

func (ar ApplicationDataRepo) GetActiveApplicationByUserIdAndJobId(UserId uint, JobId uint) (*dp.Application, error) {
	var application dp.Application
	if err := ar.DB.Database.Table("applications").Where("user_id = ? AND job_id = ? AND is_active = TRUE", UserId, JobId).First(&application).Error; err != nil {
		ar.Log.Error(err)
		return nil, err
	}

	return &application, nil
}

func (ar ApplicationDataRepo) UpdateStatus(application *dp.Application) (*dp.Application, error) {
	if err := ar.DB.Database.Model(&application).Where("user_id = ? AND job_id = ?", application.UserId, application.JobId).Update("status", application.Status).Update("is_active", application.IsActive).Error; err != nil {
		ar.Log.Error(err)
		return nil, err
	}

	return application, nil
}

func (ar ApplicationDataRepo) GetAllJobApplicationByJobId(JobId uint) ([]dp.Application, error) {
	var applications []dp.Application
	if err := ar.DB.Database.Preload("User").Table("applications").Where("job_id = ?", JobId).Find(&applications).Error; err != nil {
		ar.Log.Error(err)
		return nil, err
	}

	return applications, nil
}

func (ar ApplicationDataRepo) GetAllApplicationByUserIdAndJobId(UserId uint, JobId uint) ([]dp.Application, error) {
	var applications []dp.Application
	if err := ar.DB.Database.Table("applications").Order("created_at desc").Where("user_id = ? AND job_id = ?", UserId, JobId).Find(&applications).Error; err != nil {
		ar.Log.Error(err)
		return nil, err
	}

	return applications, nil
}
