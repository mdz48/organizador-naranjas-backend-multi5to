package entities

// StatusMessage representa un mensaje de cambio de estado para un ESP32
type StatusMessage struct {
	Esp32ID string `json:"esp32_fk"`
	Content string `json:"content"`
}

// IsValid verifica si el estado es válido
func (m *StatusMessage) IsValid() bool {
	return m.Content == "esperando" ||
		m.Content == "activo" ||
		m.Content == "desactivado"
}
