package main

import (
	"context"
	action_cable "github.com/launchboxio/action-cable"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.TODO()
	clusterId := os.Getenv("CLUSTER_ID")
	credentials := clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     os.Getenv("TOKEN_URL"),
	}
	token, err := credentials.Token(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ws, err := action_cable.New(os.Getenv("WEBSOCKET_URL"), http.Header{
		"Authorization": []string{"Bearer " + token.AccessToken},
	})

	ws.OnConnect = func() {
		log.Println("Connected to websocket")
	}
	ws.OnDisconnect = func() {
		log.Println("Websocket disconnected")
	}
	ws.OnMessage = func(message []byte) {
		log.Println(string(message))
	}
	ws.ErrorHandler = func(err error) {
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	//var wg sync.WaitGroup
	clusterSubscription := action_cable.NewSubscription(map[string]string{
		"cluster_id": clusterId,
		"channel":    "ClusterChannel",
	})
	clusterSubscription.Handler(func(event *action_cable.ActionCableEvent) {
		log.Println("subscription.handler")
		log.Println(string(event.Message))
	})

	ws.Subscribe(clusterSubscription)

	//go func() {
	//	time.Sleep(time.Second * 2)
	//	log.Println("Try a manual connection")
	//	ws.SendRaw([]byte("{\"command\": \"subscribe\",\"identifier\": \"{\\\"channel\\\":\\\"ClusterChannel\\\",\\\"cluster_id\\\":\\\"1\\\"}\"}"))
	//}()

	if err = ws.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Process completed")
	//
	//go func() {
	//	wg.Add(1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println("Oops, we're done")
	//	defer wg.Done()
	//}()

	//ws.Subscribe("ping", func(event *action_cable.ActionCableEvent) {
	//})

	//log.Println("Connecting to subscription implicitly")
	//if err := subscription.Connect(); err != nil {
	//	log.Printf("Subscription failed: %v", err)
	//} else {
	//	log.Println("Subscribed to channel")
	//}
	//
	//log.Println("Try a manual connection")
	//ws.SendRaw([]byte("{\"command\": \"subscribe\",\"identifier\": \"{\\\"channel\\\":\\\"ClusterChannel\\\",\\\"cluster_id\\\":\\\"1\\\"}\"}"))
	//wg.Wait()
}
