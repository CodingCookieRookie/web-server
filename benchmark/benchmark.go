package main

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	serverAddr  = "localhost:8080" // adjust to server address
	numClients  = 150
	numMessages = 10
	message     = "Hi this is a testing"
)

func main() {
	var wg sync.WaitGroup
	connectionSuccesses := 0
	connectionFailures := 0
	messagesSent := 0
	messagesFailed := 0

	startTime := time.Now()

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			// Create websocket connection
			u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/"}
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				fmt.Printf("Client %d: error connecting to WebSocket server: %v\n", clientID, err)
				connectionFailures++
				return
			}
			defer conn.Close()
			connectionSuccesses++

			for j := 0; j < numMessages; j++ {
				// Write message to connection
				if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					fmt.Printf("Client %d: error writing to WebSocket: %v\n", clientID, err)
					messagesFailed++
					continue
				}
				messagesSent++
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("Test completed in %s\n", duration)
	fmt.Printf("Connection successes: %d\n", connectionSuccesses)
	fmt.Printf("Connection failures: %d\n", connectionFailures)
	fmt.Printf("Messages sent: %d\n", messagesSent)
	fmt.Printf("Messages failed: %d\n", messagesFailed)
	fmt.Printf("Throughput (messages/sec): %.2f\n", float64(messagesSent)/duration.Seconds())
}
