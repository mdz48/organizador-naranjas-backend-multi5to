package messaging

import (
	"encoding/json"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain/ports"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	conn *amqp.Connection
}

func NewRabbitMQProducer(conn *amqp.Connection) *RabbitMQProducer {
	return &RabbitMQProducer{
		conn: conn,
	}
}

func (r *RabbitMQProducer) SendMessage(message ports.Message) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"State", // nombre de la cola
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	// Enlazar la cola al exchange
	err = ch.QueueBind(
		"State",      // nombre de la cola
		"api2.state", // routing key
		"amq.topic",  // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	messageSend, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return ch.Publish(
		"amq.topic", // exchange
		"api.state", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         messageSend,
		})
}
