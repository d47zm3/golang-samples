package main

import (
	"cloudy/app/app"
	"cloudy/app/router"
	"cloudy/config"
	"fmt"
	"net/http"

	dbConn "cloudy/adapter/gorm"
	lr "cloudy/util/logger"
	vr "cloudy/util/validator"
)

func main() {

	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	validator := vr.New()

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}

	if appConf.Debug {
		db.LogMode(true)
	}

	application := app.New(logger, db, validator)
	appRouter := router.New(application)
	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting Server %v", address)

	//logger.Info ("Starting Server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server Startup Failed!")
	}
}
