package config

import (
	"errors"

	"github.com/hansandika/go-job-portal-api/domain/application"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/domain/job"
	"github.com/hansandika/go-job-portal-api/domain/user"
	"github.com/hansandika/go-job-portal-api/handler"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	"github.com/hansandika/go-job-portal-api/usecase"

	"github.com/sirupsen/logrus"
)

func NewRepoContext(conf *general.SectionService) (handler.Handler, *logrus.Logger, error) {
	var handlerContext handler.Handler

	//* Init Log
	logger := infra.NewLogger(conf)
	if logger == nil {
		return handlerContext, nil, errors.New("failed to initialize logger")
	}

	// Intialize Database Connection
	database := infra.NewConnection(conf)
	if database == nil {
		return handlerContext, nil, errors.New("failed to connect to database")
	}

	// Auto Migrate
	database.Database.AutoMigrate(&user.User{})
	database.Database.AutoMigrate(&job.Job{})
	database.Database.AutoMigrate(&application.Application{})

	// Initialize Repo
	repo := repository.NewRepo(database, logger)

	// Initialize Usecase
	usecase := usecase.NewUsecase(repo, conf, database, logger)

	// Initialize Handler
	handlerContext = handler.NewHandler(usecase, conf, logger)

	return handlerContext, logger, nil
}
