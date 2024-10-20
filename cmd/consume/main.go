package main

import (
	"events/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs)
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		msg.Ack(false)
	}
}
