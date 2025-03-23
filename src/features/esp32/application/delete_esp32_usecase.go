package application

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
)

type DeleteEsp32UseCase struct {
	db ports.IEsp32
}

func NewDeleteEsp32UseCase(db ports.IEsp32) *DeleteEsp32UseCase {
	return &DeleteEsp32UseCase{
		db: db,
	}
}

func (uc *DeleteEsp32UseCase) Run(id string) error {
	return uc.db.Delete(id)
}
