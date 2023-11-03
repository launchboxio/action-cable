package mux_socket

// WebsocketMessage represents the parsed payload
// returned from the stream
type WebsocketMessage struct {
	Channel string
	Event   Event[any]
}

type Event[C any] struct {
}
