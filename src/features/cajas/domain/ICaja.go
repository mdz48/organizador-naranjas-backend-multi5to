package domain

type ICaja interface {
	GetAll() ([]Caja, error)
	GetById(id int) (Caja, error)
	GetByDescripcion(descripcion string) (Caja, error)
	Create(caja Caja) (Caja, error)
	Update(caja Caja) (Caja, error)
	Delete(id int) error
	GetByLote(lote int) ([]Caja, error)
}
