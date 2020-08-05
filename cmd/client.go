package main

import (
	"log"
	"os"
	"time"

	"github.com/kowiste/gos7"
)

func main() {
	//Testing the connection
	handler := gos7.NewTCPClientHandler("127.0.0.1", 5, 2)
	handler.Timeout = 5 * time.Millisecond
	handler.IdleTimeout = 500 * time.Millisecond
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
