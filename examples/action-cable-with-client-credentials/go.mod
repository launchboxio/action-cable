module github.com/launchboxio/mux_socket/examples/action-cable-with-client-credentials

go 1.21.1

replace github.com/launchboxio/action-cable => ../../

require (
	github.com/launchboxio/action-cable v0.0.0-20231102203233-655e55a949c1
	golang.org/x/oauth2 v0.13.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
