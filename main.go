package main

import (
	"fmt"
	"runtime"

	"github.com/hmendezp/godesde0/ejercicios"
	"github.com/hmendezp/godesde0/variables"
)

func main() {
	fmt.Println("Hola Mundo Chingón!!")
	variables.MostrarEnteros()
	variables.MostrarRestoVars()
	estado, textRet := variables.Convertir2Texto(1001)
	fmt.Println(estado, textRet)

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Tenemos un pingüino")
	} else {
		fmt.Println("Tenemos otro sistema distinto!", os)
	}

	entero, cadena := ejercicios.Excer01("77")
	fmt.Println(entero, cadena)
}
