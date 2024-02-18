package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/handler"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *gin.Engine, handler handler.Handler, conf *general.SectionService, log *logrus.Logger) {

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Job Portal API",
		})
	})

	// Api v1
	api := app.Group(conf.App.Endpoint)

	// Endpoint
	RouteUser(api, handler, conf, log)
	RouteJob(api, handler, conf, log)
	RouteApplication(api, handler, conf, log)
}
