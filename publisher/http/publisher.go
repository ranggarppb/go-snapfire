package http

import (
	"go-snapfire/types"
	"io"
	"net/http"
	"strings"
)

type httpPublisher struct {
	listenAddr string

	produceChannel chan<- types.Message
}

func NewHttpPublisher(addr string) *httpPublisher {
	return &httpPublisher{
		listenAddr: addr,
	}
}

func (p *httpPublisher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)

	switch {
	case r.Method == http.MethodPost:
		body, _ := io.ReadAll(r.Body)
		p.produceChannel <- types.Message{
			Topic: parts[0],
			Data:  body,
		}

		defer r.Body.Close()

		w.WriteHeader(http.StatusOK)
	}
}

func (p *httpPublisher) Start(produceChan chan<- types.Message) error {
	p.produceChannel = produceChan
	return http.ListenAndServe(p.listenAddr, p)
}
