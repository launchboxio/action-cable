package mux_socket

type Adapter interface {
	// Marshal converts a tangible event into a
	// websocket compliant message
	Marshal(message WebsocketMessage) ([]byte, error)

	// Unmarshal translates a websocket message
	// into an Event object
	Unmarshal([]byte) (WebsocketMessage, error)
}
