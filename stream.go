package mux_socket

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"time"
)

type Stream struct {
	Adapter Adapter
	Url     string
	Header  http.Header

	NotFoundHandler HandlerFunc
	ErrorHandler    func(error)
	OnDisconnect    func()
	OnConnect       func()

	subscriptions map[string][]HandlerFunc

	send chan []byte

	connected  bool
	connection *websocket.Conn
}

type HandlerFunc func(event Event)

type Route struct {
	Channel string
	Handler HandlerFunc
}

func New(url string, adapter Adapter, header http.Header) (*Stream, error) {
	return &Stream{
		Adapter: adapter,
		Url:     url,
		Header:  header,
	}, nil
}

// Handler registers a handler function for a given channel
func (s *Stream) Handler(channel string, handler HandlerFunc) {
	if _, ok := s.subscriptions[channel]; !ok {
		s.subscriptions[channel] = []HandlerFunc{}
	}
	s.subscriptions[channel] = append(s.subscriptions[channel], handler)
}

// Connect connects to the websocket stream, and
// dispatches any received messages
func (s *Stream) Connect(ctx context.Context) error {
	u, err := url.Parse(s.Url)
	if err != nil {
		return err
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), s.Header)
	if err != nil {
		return err
	}
	done := make(chan struct{})

	s.onConnect(c)
	defer c.Close()

	go func() {
		defer close(done)
		s.listen()
	}()

	for {
		select {
		case <-done:
			return nil
		case <-s.send:

		case <-ctx.Done():
			return s.close(done)
		}
	}
}

func (s *Stream) onConnect(connection *websocket.Conn) {
	s.connection = connection
	s.connected = true
	s.send = make(chan []byte)

	// Notify anyOnConnect handlers
	if s.OnConnect != nil {
		s.OnConnect()
	}
}

func (s *Stream) close(done chan struct{}) error {
	err := s.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	s.OnDisconnect()
	select {
	case <-done:
	case <-time.After(time.Second):
	}
	return nil
}

func (s *Stream) listen() {
	for {
		_, message, err := s.connection.ReadMessage()
		if err != nil {
			s.ErrorHandler(err)
			continue
		}
		if err = s.process(message); err != nil {
			s.ErrorHandler(err)
			continue
		}
	}
}

// process parses an incoming message using the configured
// adapter, and then passes it along for dispatching
func (s *Stream) process(message []byte) error {
	// Parse the message using our adapter, t
	event, err := s.Adapter.Unmarshal(message)
	if err != nil {
		return err
	}

	return s.dispatch(event)
}

// dispatch takes a parsed event, and passes it off
// to all of the registered handlers for its channel
func (s *Stream) dispatch(message WebsocketMessage) error {
	handlers, found := s.subscriptions[event.Channel]
	if !found {
		s.NotFoundHandler(event)
		return errors.New(fmt.Sprintf("Channel %s not found for event", event.Channel))
	} else {
		for _, handler := range handlers {
			// TODO: Support middleware
			handler(event)
		}
		return nil
	}
}

func (s *Stream) IsConnected() bool {
	return s.connected
}

func (s *Stream) Send(event Event) error {
	bytes, err := s.Adapter.Marshal(event)
	if err != nil {
		s.ErrorHandler(err)
		return err
	}

	s.send <- bytes
	return nil
}
