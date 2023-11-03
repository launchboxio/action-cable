package adapter

import (
	"encoding/json"
	"github.com/launchboxio/mux-socket"
)

type ActionCable struct {
	Identifier map[string]string `json:"identifier"`
}

type ActionCableEvent[C any] struct {
	Command    string `json:"command"`
	Data       C      `json:"data,omitempty"`
	Identifier string `json:"identifier"`
}

func (ac *ActionCable) Marshal(event mux_socket.WebsocketMessage) ([]byte, error) {
	identifier, err := json.Marshal(ac.Identifier)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (ac *ActionCable) Unmarshal(channel string, data []byte) (*mux_socket.WebsocketMessage, error) {
	result := &mux_socket.WebsocketMessage{
		Channel: channel,
		Event:
	}
}
