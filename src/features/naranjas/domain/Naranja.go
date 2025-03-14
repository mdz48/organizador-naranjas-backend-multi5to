package domain
import "time"

type Naranja struct {
	ID      int     	`json:"id"`
	Peso 	string  	`json:"peso"`
	Tamaño  float32 	`json:"tamaño"`
	Color   float32 	`json:"color"`
	Hora	time.Time 	`json: "hora"`
	CajaFK	time.Time 	`json: "caja_fk`
}
