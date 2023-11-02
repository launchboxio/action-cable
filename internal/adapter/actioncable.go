package adapter

import (
	"github.com/launchboxio/mux-socket"
)

type ActionCable struct {
}

func (ac *ActionCable) Marshal(*mux_socket.Event) ([]byte, error) {
	return nil, nil
}

func (ac *ActionCable) Unmarshal([]byte) (*mux_socket.Event, error) {
	return nil, nil
}
