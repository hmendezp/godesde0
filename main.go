package main

import (
	"fmt"

	"github.com/hmendezp/godesde0/variables"
)

func main() {
	fmt.Println("Hola Mundo Chingón!!")
	variables.MostrarEnteros()
	variables.MostrarRestoVars()
	estado, textRet := variables.Convertir2Texto(1001)
	fmt.Println(estado, textRet)
}
