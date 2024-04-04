package main

import (
	"brequet/goethe/internal/app/api/dispatcher"
	"brequet/goethe/internal/app/api/handler"
	"brequet/goethe/internal/webapp"
	"brequet/goethe/pkg/hello"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	logger := log.New(os.Stdout, "", log.LstdFlags)

	webAppHandler, err := webapp.NewWebAppHandler(logger)
	if err != nil {
		logger.Fatal("error while creating web app handler:", err)
	}
	apiHandler := handler.NewExampleApiHandler(httpClient, hello.NewHelloService())

	router := dispatcher.NewRouter(logger, "localhost", 3000, webAppHandler, apiHandler)
	if err := router.Dispatch(); err != nil {
		logger.Fatal("error while serving the application:", err)
	}
}
