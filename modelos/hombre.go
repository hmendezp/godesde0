package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

// Lo siguiente es similar a los Getters y Setters de Java
func (h *Hombre) Respirar()         { h.respirando = true }
func (h *Hombre) Comer()            { h.comiendo = true }
func (h *Hombre) Pensar()           { h.pensando = true }
func (h *Hombre) SetVida(vida bool) { h.vivo = vida }
func (h *Hombre) Sexo() string      { return "Hombre" }
func (h *Hombre) EstaVivo() bool    { return h.vivo }
