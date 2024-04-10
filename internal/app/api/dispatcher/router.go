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
	if err := http.ListenAndServe(serverAddrPort, corsMiddleware(mux)); err != nil {
		return err
	}

	return nil
}

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // or "*" for any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if it's a preflight request
		if r.Method == "OPTIONS" {
			// If so, handle the preflight and stop the chain
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
