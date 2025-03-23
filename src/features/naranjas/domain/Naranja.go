package domain

import "time"

type Naranja struct {
    ID     int       `json:"id"`
    Peso   float32   `json:"peso"`
    Tamaño string    `json:"tamano"`
    Color  string    `json:"color"`
    Hora   time.Time `json:"hora"`
    CajaFK int       `json:"caja_fk"`
    Esp32FK string   `json:"esp32_fk"`
}