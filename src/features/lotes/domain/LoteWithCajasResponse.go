package domain

import (
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type LoteWithCajasResponse struct {
	Lote  Lote              `json:"lote"`
	Cajas []cajaDomain.Caja `json:"cajas"`
}
