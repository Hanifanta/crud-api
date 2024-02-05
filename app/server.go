package app

import (
	"crud-api/utils"
	"net/http"
)

func Serve(config utils.Config) error {
	router := NewRouter(config)
	server := http.Server{
		Addr:    config.ServerAddress,
		Handler: router,
	}

	return server.ListenAndServe()
}
