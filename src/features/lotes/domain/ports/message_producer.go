package ports

// Message representa el mensaje que se enviará a RabbitMQ
type Message struct {
	Esp32FK string `json:"esp32_fk"`
	Content string `json:"content"`
}

// IMessageProducer define el puerto para enviar mensajes
type IMessageProducer interface {
	SendMessage(message Message) error
}
