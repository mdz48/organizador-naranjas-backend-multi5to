package domain

type Lote struct {
	ID            int    `json:"id"`
	Fecha         string `json:"fecha"`
	Observaciones string `json:"observaciones"`
	Estado        string `json:"estado"` 
}
