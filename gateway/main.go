package main

import (
	"log"
    "net/http"
)

func main() {
    http.HandleFunc("/ws", wsHandler)

    log.Println("Gateway started on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}