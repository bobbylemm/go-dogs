package server

import (
	"fmt"
	"go-dogs/config"
	"net/http"
)

func StartServer() error {
	appConfig := config.GetConfig()
	router := NewRouter(*appConfig)
	router.initializeRoutes()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.AppPort), router); err != nil {
		return err
	}
	return nil
}