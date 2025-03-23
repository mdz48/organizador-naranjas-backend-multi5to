package domain

type IProducer interface {
	PublishNaranja(naranja Naranja) error
}