package ports

// IEsp32Producer define la interfaz para enviar mensajes de estado de ESP32
type IEsp32Producer interface {
	PublishStatusChange(esp32ID string, status string) error
}
