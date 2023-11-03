package adapter

import mux_socket "github.com/launchboxio/mux-socket"

type DefaultAdapter struct {
}

func (da *DefaultAdapter) Marshal(mux_socket.Event) ([]byte, error) {
	return nil, nil
}

func (da *DefaultAdapter) Unmarshal([]byte) (mux_socket.Event, error) {
	return mux_socket.Event{}, nil
}
