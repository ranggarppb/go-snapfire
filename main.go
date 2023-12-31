package main

import (
	httpPublisher "go-snapfire/publisher/http"
	"go-snapfire/types"
	"log"
)

func main() {
	topicStore := map[string]types.Storage{}
	publisher := httpPublisher.NewHttpPublisher(":3000")
	publishers := []types.Publisher{publisher}

	server, err := NewServer(topicStore, publishers)

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
