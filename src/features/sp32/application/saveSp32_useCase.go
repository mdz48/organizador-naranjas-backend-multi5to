package application

import (
	"log"
	"organizador-naranjas-backend-multi5to/src/features/sp32/domain/entities"
	"organizador-naranjas-backend-multi5to/src/features/sp32/domain/ports"
)

type SaveSp32UseCase struct {
	sp32_repository ports.ISp32
}

func NewSaveSp32(sp32_repository ports.ISp32) *SaveSp32UseCase {
	return &SaveSp32UseCase{
		sp32_repository: sp32_repository,
	}
}

func (uc *SaveSp32UseCase) Run(sp32 *entities.Sp32) (*entities.Sp32, error) {
	sp32, err := uc.sp32_repository.Save(sp32)

	if err != nil {
		log.Printf("error to save sp32")
		return &entities.Sp32{}, err
	}

	return sp32, nil
}