package main

import (
	"fmt"
	"runtime"

	"github.com/hmendezp/godesde0/ejercicios"
	"github.com/hmendezp/godesde0/files"
	"github.com/hmendezp/godesde0/iteraciones"
	"github.com/hmendezp/godesde0/teclado"
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
		fmt.Println("Tenemos otro sistema distinto!")
	}

	switch os2 := runtime.GOOS; os2 {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Println(os2)
	}

	entero, cadena := ejercicios.Excer01("-77")
	fmt.Println(entero, cadena)

	// teclado.IngresNumber()
	teclado.Nada()

	iteraciones.Iterar()

	// ejercicios.DespliegaTabla()
	// ejercicios.MostrarTablaDel(8)
	numero := 8
	fmt.Print(ejercicios.ArmaTablaTexto(numero))
	// files.EscribeArchTabla(numero)
	// files.AgregaTabla(numero)
	files.LeoArchivo()
}
