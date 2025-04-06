package adapters

import (
	"encoding/json"
	"log"
	"organizador-naranjas-backend-multi5to/src/core"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQProducer implementa la interfaz IEsp32Producer
type RabbitMQProducer struct {
	rabbitMQ *core.RabbitMQ
}

// NewRabbitMQProducer crea una nueva instancia de RabbitMQProducer
func NewRabbitMQProducer(rabbitMQ *core.RabbitMQ) *RabbitMQProducer {
	return &RabbitMQProducer{
		rabbitMQ: rabbitMQ,
	}
}

// PublishStatusChange publica un mensaje de cambio de estado en RabbitMQ
func (r *RabbitMQProducer) PublishStatusChange(esp32ID string, status string) error {
	// Si no hay conexión a RabbitMQ, logueamos y retornamos sin error
	if r.rabbitMQ == nil {
		log.Printf("RabbitMQ no está disponible, no se puede enviar mensaje de estado para ESP32 %s", esp32ID)
		return nil
	}

	// Obtener un canal de RabbitMQ
	ch, err := r.rabbitMQ.Conn.Channel()
	if err != nil {
		log.Printf("Error al obtener canal RabbitMQ: %v", err)
		return err
	}
	defer ch.Close()

	// Declarar la cola
	_, err = ch.QueueDeclare(
		"State", // nombre de la cola
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("Error al declarar cola RabbitMQ: %v", err)
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
		log.Printf("Error al enlazar cola RabbitMQ: %v", err)
		return err
	}

	// Crear el mensaje
	message := entities.StatusMessage{
		Esp32ID: esp32ID,
		Content: status,
	}

	// Serializar el mensaje a JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error al serializar mensaje: %v", err)
		return err
	}

	// Publicar el mensaje
	err = ch.Publish(
		"amq.topic", // exchange
		"api.state", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         messageBytes,
		})
	if err != nil {
		log.Printf("Error al publicar mensaje: %v", err)
		return err
	}

	log.Printf("Mensaje de estado enviado para ESP32 %s: %s", esp32ID, status)
	return nil
}
