package main

import (
	"net"
	"log"
)

type Producer struct {
	message string
}

func NewProducer(message string) Producer {
	return Producer{
		message: message,
	}
}

func (p Producer) SendMessage(addr string) error {

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte(p.message))
	return err
}
