package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/handler"
	"github.com/hansandika/go-job-portal-api/middleware"
	"github.com/sirupsen/logrus"
)

func RouteApplication(api *gin.RouterGroup, handler handler.Handler, conf *general.SectionService, log *logrus.Logger) {
	api.POST("/application/create/:jobId", middleware.JWTMiddleware(conf, log), handler.Application.Application.CreateApplication)
	api.PATCH("/application/update", middleware.JWTMiddleware(conf, log), middleware.EmployerMiddleware(conf, log), handler.Application.Application.UpdateApplication)
	api.GET("/application/:jobId", middleware.JWTMiddleware(conf, log), handler.Application.Application.GetLoggedUserApplication)
	api.GET("/application/:jobId/:email", middleware.JWTMiddleware(conf, log), middleware.EmployerMiddleware(conf, log), handler.Application.Application.GetJobApplicationsByUserEmailAndJobId)
	api.GET("/application/list/:jobId", middleware.JWTMiddleware(conf, log), middleware.EmployerMiddleware(conf, log), handler.Application.Application.GetJobApplicationsByJobId)
}
