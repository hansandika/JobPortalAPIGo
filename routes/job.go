package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/handler"
	"github.com/hansandika/go-job-portal-api/middleware"
	"github.com/sirupsen/logrus"
)

func RouteJob(api *gin.RouterGroup, handler handler.Handler, conf *general.SectionService, log *logrus.Logger) {
	api.POST("/job/create", middleware.JWTMiddleware(conf, log), middleware.EmployerMiddleware(conf, log), handler.Job.Job.CreateJob)
	api.GET("/job/list", middleware.JWTMiddleware(conf, log), handler.Job.Job.ListJobs)
}
