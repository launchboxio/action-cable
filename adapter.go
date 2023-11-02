package mux_socket

type Adapter interface {
	// Marshal converts a tangible event into a
	// websocket compliant message
	Marshal(*Event) ([]byte, error)

	// Unmarshal translates a websocket message
	// into an Event object
	Unmarshal([]byte) (*Event, error)
}

type DefaultAdapter struct {
}

func (da *DefaultAdapter) Marshal(*Event) ([]byte, error) {
	return nil, nil
}

func (da *DefaultAdapter) Unmarshal([]byte) (*Event, error) {
	return nil, nil
}
