package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/handler"
	"github.com/hansandika/go-job-portal-api/middleware"
	"github.com/sirupsen/logrus"
)

func RouteUser(api *gin.RouterGroup, handler handler.Handler, conf *general.SectionService, log *logrus.Logger) {
	api.POST("/user/login", handler.User.User.Login)
	api.POST("/user/register", handler.User.User.Register)
	api.POST("/user/logout", middleware.JWTMiddleware(conf, log), handler.User.User.Logout)
}
