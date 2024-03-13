package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CodingCookieRookie/web-server/generator"
	"github.com/CodingCookieRookie/web-server/log"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{}
var host = "localhost"
var port = "8080"

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

func loadServerAddr() {
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("err upgrading connection to websocket: %v", err)
		return
	}
	defer c.Close()

	for {
		t, b, err := c.ReadMessage()
		if err != nil {
			log.Errorf("error reading message: %v", err)
			break
		}
		log.Debugf("message type: %v", t)
		switch t {
		case websocket.TextMessage:
			clientMsg := string(b)
			log.Infof("Received text message: %s", clientMsg)
		case websocket.BinaryMessage:
			log.Infof("Received binary message: %v", b)
		case websocket.CloseMessage:
			log.Infof("Received close message")
			return
		case websocket.PingMessage:
			log.Infof("Received ping message")
		case websocket.PongMessage:
			log.Infof("Received pong message")
		default:
			log.Infof("Received unhandled message type: %d", t)
		}

		uniqueBigInt := generator.GenerateUniqueBigInt()
		err = c.WriteMessage(websocket.TextMessage, []byte(uniqueBigInt.String()))
		if err != nil {
			log.Errorf("error writing response to websocket: %v", err)
			break
		}
	}
}

func main() {
	loadEnv()
	loadServerAddr()
	log.InitLogger()

	http.HandleFunc("/", handleConnection)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil); err != nil {
		log.Panicf("error listening and serving: %v", err)
	}
}
