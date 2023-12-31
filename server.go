package main

import (
	inmemorystore "go-snapfire/storage/in_memory_store"
	"go-snapfire/types"
)

type server struct {
	TopicStore map[string]types.Storage
	Publishers []types.Publisher

	produceChannel chan types.Message
	quitChannel    chan struct{}
}

func NewServer(topicStore map[string]types.Storage, publishers []types.Publisher) (*server, error) {
	return &server{
		TopicStore: topicStore,
		Publishers: publishers,

		produceChannel: make(chan types.Message),
		quitChannel:    make(chan struct{}),
	}, nil
}

func (s *server) Start() error {

	for _, publisher := range s.Publishers {
		go func(p types.Publisher) {
			p.Start(s.produceChannel)
		}(publisher)
	}

	for {
		select {
		case <-s.quitChannel:
			return nil
		case msg := <-s.produceChannel:
			topic := msg.Topic
			message := msg.Data

			topicStore := s.getOrCreateStoreForTopic(topic)
			topicStore.Push(message)
		}
	}
}

func (s *server) getOrCreateStoreForTopic(topic string) types.Storage {
	if _, ok := s.TopicStore[topic]; !ok {
		store := inmemorystore.NewInMemoryStore()
		s.TopicStore[topic] = store
	}

	return s.TopicStore[topic]
}
