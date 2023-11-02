package interfaces

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	SetVida(bool)
	EstaVivo() bool
}
