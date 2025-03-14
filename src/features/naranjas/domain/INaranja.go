package domain

type INaranja interface {
	Create(naranja Naranja) (Naranja, error)
	GetById(id int) (Naranja, error)
	GetByCaja(cajaId int) ([]Naranja, error)
	GetAll() ([]Naranja, error)
	Update(naranja Naranja) (Naranja, error)
	Delete(id int) error
}
