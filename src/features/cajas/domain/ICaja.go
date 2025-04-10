package domain

type ICaja interface {
	GetAll() ([]Caja, error)
	GetById(id int) (Caja, error)
	GetByDescripcion(descripcion string) (Caja, error)
	Create(caja Caja) (Caja, error)
	Update(caja Caja) (Caja, error)
	Delete(id int) error
	GetByLote(loteId int) ([]Caja, error)
	FindByEsp32AndState(esp32Id string, state string) (Caja, error)
	FindByEsp32StateAndDescription(esp32Id string, state string, description string) (Caja, error)
	UpdateStatusByLoteId(loteId int, estado string) error
	GetTop3ByLote(loteId int) ([]Caja, error)
	GetLotesByEsp32(esp32Id string, estado string) ([]int, error)
}
