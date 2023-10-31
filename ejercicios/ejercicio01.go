package ejercicios

import (
	"fmt"
	"strconv"
)

func Excer01(param1 string) (int, string) {
	var valorRet string
	i, erroneo := strconv.Atoi(param1)
	if erroneo != nil {
		fmt.Println("Ha ocurrido un error durante la conversión de cadena a entero con ", param1)
		panic(erroneo)
	}
	if i > 100 {
		valorRet = "Es mayor a 100"
	} else {
		valorRet = "Es menor a 100"
	}
	return i, valorRet

}
