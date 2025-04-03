package domain

type LoteWithCajas struct {
	Lote Lote `json:"lote"`
	Esp32FK string `json:"esp32_fk"`
}