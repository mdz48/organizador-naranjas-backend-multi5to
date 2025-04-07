package core

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitConn     *amqp.Connection
	rabbitConnOnce sync.Once
)

// GetRabbitMQConnection returns a singleton RabbitMQ connection
func GetRabbitMQConnection() *amqp.Connection {
	rabbitConnOnce.Do(func() {
		conn, err := amqp.Dial("amqp://user1:lsrgil@54.175.42.193:5672/")
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}
		rabbitConn = conn
	})
	return rabbitConn
}

// CloseRabbitMQConnection closes the RabbitMQ connection
func CloseRabbitMQConnection() {
	if rabbitConn != nil {
		rabbitConn.Close()
		rabbitConn = nil
	}
}
