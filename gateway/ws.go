package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ClientType string

const (
	UserClient ClientType = "user"
	PrinterClient ClientType = "priter"
)

type Client struct {
	Conn *websocket.Conn
	Type ClientType
	PrinterID string
	UserID string
}

type Message struct {
	Type string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}


var (
	users = make(map[string]*Client)
	priters = make(map[string]*Client)
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{Conn: conn}

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err  != nil{
			log.Println("read error:", err)
			return
		}

		switch msg.Type {
		case "ayth":
			handleUserAuth(client, msg.Payload)
		case "register_printer":
			handlePrinterRegister(client, msg.Payload)
		case "telemetry":
			handleTelemetry(client, msg.Payload)
		default:
			log.Println("unknown message type:", msg.Type)
		}
	}
}

func handleUserAuth(c *Client, payload json.RawMessage) {
	c.Type = UserClient
	c.UserID = "user-1"
	users[c.UserID] = c
	log.Println("user authentificated:", c.UserID)
}
func handlePrinterRegister(c *Client, payload json.RawMessage) {
	
	
	c.Type = PrinterClient
	c.PrinterID = p.PrinterID
	priters[c.PrinterID] = c
	log.Println("printer registered:", c.PrinterID)
}
func handleTelemetry(c *Client, payload json.RawMessage) {
	log.Println("telemetry from", c.PrinterID, string(payload))
}