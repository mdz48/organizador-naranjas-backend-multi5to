package domain

type ILote interface {
	Create(lote Lote) (Lote, error)
	GetAll() ([]Lote, error)
	GetById(id int) (Lote, error)
	Delete(id int) error
	Update(id int, lote Lote) (Lote, error)
	GetByDate(date string) ([]Lote, error)
	GetByUserId(userId int) ([]Lote, error)
}
