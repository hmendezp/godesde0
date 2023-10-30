package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func MostrarRestoVars() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func Convertir2Texto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
