package ports

import "organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"

type IEsp32 interface {
	Save(esp32 *entities.Esp32) (*entities.Esp32, error)
	GetByPropietario(id int) ([]entities.Esp32, error)
	Delete(id string) error
}
