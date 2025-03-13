package domain

type Caja struct {
	ID          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	PesoTotal   float32 `json:"peso_total"`
	Precio      float32 `json:"precio"`
	LoteFK      int     `json:"lote_fk"`
	EncargadoFK int     `json:"encargado_fk"`
}
