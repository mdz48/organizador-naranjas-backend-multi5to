package domain

type INaranja interface {
	GetAll() ([]Naranja, error)
	GetById(id int) (Naranja, error)
	GetByCaja(cajaId int) ([]Naranja, error)
	Create(naranja Naranja) (Naranja, error)
	Update(naranja Naranja) (Naranja, error)
	Delete(id int) error
}
