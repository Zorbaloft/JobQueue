package main

import (
	"log"
	"time"
)

func main() {
	p := NewProducer("Hello World")
	s := NewServer(":8080")

	if err := p.SendMessage("localhost:8080"); err != nil {
		log.Fatal(err)
	}
	
	time.Sleep(100 * time.Millisecond) // dá tempo ao servidor para processar
	s.Stop()
}