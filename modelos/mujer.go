package modelos

type Mujer struct {
	Hombre
}

// Lo siguiente es similar a los Getters y Setters de Java
func (m *Mujer) Respirar()         { m.respirando = true }
func (m *Mujer) Comer()            { m.comiendo = true }
func (m *Mujer) Pensar()           { m.pensando = true }
func (m *Mujer) SetVida(vida bool) { m.vivo = vida }
func (m *Mujer) Sexo() string      { return "Mujer" }
func (m *Mujer) EstaVivo() bool    { return m.vivo }
