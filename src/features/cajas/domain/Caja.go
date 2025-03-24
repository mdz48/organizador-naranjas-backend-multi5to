package domain

import "time"

type Caja struct {
    ID          int       `json:"id"`
    Descripcion string    `json:"descripcion"`
    PesoTotal   float32   `json:"peso_total"`
    Precio      float32   `json:"precio"`
    HoraInicio  time.Time `json:"hora_inicio"`
    HoraFin     time.Time `json:"hora_fin"`
    LoteFK      int       `json:"lote_fk"`
    EncargadoFK int       `json:"encargado_fk"`
    Cantidad    int       `json:"cantidad"`
    Estado      string    `json:"estado"`    
    Esp32FK     string    `json:"esp32_fk"`   
}