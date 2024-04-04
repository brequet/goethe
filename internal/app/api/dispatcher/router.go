package dispatcher

import (
	"brequet/goethe/internal/app/api/handler"
	"brequet/goethe/internal/webapp"
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	logger            *log.Logger
	serverAddr        string
	serverPort        int
	webAppHandler     *webapp.WebAppHandler
	exampleApiHandler *handler.ExampleApiHandler
}

func NewRouter(
	logger *log.Logger,
	serverAddr string,
	serverPort int,
	webAppHandler *webapp.WebAppHandler,
	exampleApiHandler *handler.ExampleApiHandler,
) *Router {
	return &Router{
		logger:            logger,
		serverAddr:        serverAddr,
		serverPort:        serverPort,
		webAppHandler:     webAppHandler,
		exampleApiHandler: exampleApiHandler,
	}
}

func (r Router) Dispatch() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.webAppHandler.ServeHTTP)

	mux.HandleFunc("/api/hello", r.exampleApiHandler.GetHello)
	mux.HandleFunc("/api/goodbye", r.exampleApiHandler.GetGoodbye)

	serverAddrPort := fmt.Sprintf("%s:%d", r.serverAddr, r.serverPort)
	r.logger.Printf("Serving web app on %s...\n", serverAddrPort)
	if err := http.ListenAndServe(serverAddrPort, mux); err != nil {
		return err
	}

	return nil
}
