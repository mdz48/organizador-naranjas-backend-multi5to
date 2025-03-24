package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type FindActiveCajaByEsp32UseCase struct {
	cajaRepository domain.ICaja
}

func NewFindActiveCajaByEsp32UseCase(cajaRepository domain.ICaja) *FindActiveCajaByEsp32UseCase {
	return &FindActiveCajaByEsp32UseCase{
		cajaRepository: cajaRepository,
	}
}

func (c *FindActiveCajaByEsp32UseCase) Execute(esp32Id string, color string) (domain.Caja, error) {
	return c.cajaRepository.FindByEsp32StateAndDescription(esp32Id, "CARGANDO", color)
}
