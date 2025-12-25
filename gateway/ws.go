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
	PrinterClient ClientType = "printer"
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
	printers = make(map[string]*Client)
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
	
	defer func() {
		if client.Type == PrinterClient {
			delete(printers, client.PrinterID)
			log.Println("printer disconnectes:", client.PrinterID)
		}

		if client.Type == UserClient {
			delete(users, client.UserID)
			log.Println("user disconnected:", client.UserID)
		}
	}()

	log.Println("new ws connection")

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err  != nil{
			log.Println("read error:", err)
			return
		}

		switch msg.Type {
		case "auth":
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
	var p struct {
		PrinterID string `json:"printer_id"`
	}
	
	if err := json.Unmarshal(payload, &p); err != nil {
		log.Println("invalid printer payload:", err)
		return
	}

	if p.PrinterID == "" {
		log.Println("empty printer_id in register_printer")
		return
	}

	c.Type = PrinterClient
	c.PrinterID = p.PrinterID
	printers[c.PrinterID] = c

	send(c,"printer_registered", map[string]string{
		"printer_id": c.PrinterID,
	})

	log.Println("printer registered:", c.PrinterID)
}
func handleTelemetry(c *Client, payload json.RawMessage) {
	if c.Type != PrinterClient {
		log.Println("telemetry from non-printer client")
		return
	}

	log.Println("telemetry from", c.PrinterID, string(payload))
}

func send(c *Client, msgType string, payload any){
	msg := map[string]any{
		"type": msgType,
		"payload": payload,
	}

	if err := c.Conn.WriteJSON(msg); err != nil {
		log.Println("send error:", err)
	}
}