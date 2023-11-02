package funciones

import "fmt"

var indice = 0

func Exponencia(valor int) {
	if valor > 1000000000 {
		return
	} else {
		indice++
	}

	fmt.Println("2 ^ ", indice, " = ", valor)
	Exponencia(valor * 2)
}
