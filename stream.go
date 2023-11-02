package mux_socket

import "context"

type Stream struct {
	Adapter Adapter

	NotFoundHandler HandlerFunc
	OnDisconnect    HandlerFunc
	OnConnect       HandlerFunc

	subscriptions map[string][]HandlerFunc
}

type HandlerFunc func(event Event)

type Route struct {
	Channel string
	Handler HandlerFunc
}

func New(adapter Adapter) (*Stream, error) {
	return &Stream{
		Adapter: adapter,
	}, nil
}

// Handler registers a handler function for a given channel
func (s *Stream) Handler(channel string, handler HandlerFunc) {
	if _, ok := s.subscriptions[channel]; !ok {
		s.subscriptions[channel] = []HandlerFunc{}
	}
	s.subscriptions[channel] = append(s.subscriptions[channel], handler)
}

func (s *Stream) Listen(ctx context.Context) {

}
