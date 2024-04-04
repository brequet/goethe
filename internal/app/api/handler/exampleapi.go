package handler

import (
	"brequet/goethe/pkg/hello"
	"net/http"
)

type ExampleApiHandler struct {
	httpClient   *http.Client
	helloService *hello.HelloService
}

func NewExampleApiHandler(httpClient *http.Client, helloService *hello.HelloService) *ExampleApiHandler {
	return &ExampleApiHandler{
		httpClient:   httpClient,
		helloService: helloService,
	}
}

func (h *ExampleApiHandler) GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.helloService.SayHello()))
}

func (h *ExampleApiHandler) GetGoodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.helloService.SayGoodBye()))
}
