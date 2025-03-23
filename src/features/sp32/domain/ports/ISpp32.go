package ports

import "organizador-naranjas-backend-multi5to/src/features/sp32/domain/entities"

type ISp32 interface {
	Save(sp32 *entities.Sp32) (*entities.Sp32, error)
}