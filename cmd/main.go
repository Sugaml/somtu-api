package main

/*
*************************************************************************

Author: Babulal Tamang
Purpose: LMS
Model Name:
Date: 15th June 2025
Additional Notes:
Refactored: 7th Feb 2025

****************************************************************************
*/
import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/adaptor/config"
	"github.com/sugaml/lms-api/internal/adaptor/http"
	"github.com/sugaml/lms-api/internal/adaptor/storage/postgres"
	"github.com/sugaml/lms-api/internal/adaptor/storage/postgres/repository"
	"github.com/sugaml/lms-api/internal/adaptor/storage/uploader"
	"github.com/sugaml/lms-api/internal/core/auth"
	"github.com/sugaml/lms-api/internal/core/service"
)

// @title						LMS API
// @version						1.0
// @description					This is a simple RESTful Service API for LMS Management
// @basePath					/api/v1/lms
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 							Header
// @name 						Authorization
func main() {
	// Load environment variables
	config, err := config.LoadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading environment variables")
	}
	logrus.Infof("Starting the application %s in %s mode", config.APP_NAME, config.APP_ENV)

	// Init database
	db, err := postgres.NewDB(config)
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing database connection")
	}

	repo := repository.NewRepository(db)
	tokenMaker, err := auth.NewJWTMaker(config.JWT_SECRET)
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing token maker")
	}
	svc := service.NewService(repo, tokenMaker)
	uploader, err := uploader.GetUploader()
	handler := http.NewHandler(svc, config, tokenMaker, uploader)

	// Init router
	router, err := http.NewRouter(
		config,
		*handler,
	)
	if err != nil {
		logrus.WithError(err).Fatal("Error initializing router")
	}

	listenAddr := fmt.Sprintf(":%s", config.APP_PORT)
	logrus.Infof("Starting the HTTP server in port %s", listenAddr)
	if err := router.Serve(listenAddr); err != nil {
		logrus.WithError(err).Fatal("Error starting the HTTP server")
	}
}
