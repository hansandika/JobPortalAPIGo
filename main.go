package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/global/config"
	"github.com/hansandika/go-job-portal-api/middleware"
	"github.com/hansandika/go-job-portal-api/routes"
)

func main() {
	// Initialize Gin App
	app := gin.Default()

	conf, err := config.GetCoreConfig()
	if err != nil {
		panic(err)
	}

	// Initiate Handler
	handler, logger, err := config.NewRepoContext(conf)
	if err != nil {
		panic(err)
	}

	// Initial Routes
	routes.SetupRoutes(app, handler, conf, logger)

	// Use general middleware
	middleware.GeneralMiddleware(app, conf, logger)

	// Run Application
	logger.Fatal(app.Run(fmt.Sprintf(":%s", conf.App.Port)))
}
