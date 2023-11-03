package main

import (
	"fmt"
	"runtime"

	"github.com/hmendezp/godesde0/arreglos_slices"
	d "github.com/hmendezp/godesde0/defer_panic"
	e "github.com/hmendezp/godesde0/ejer_interfaces"
	"github.com/hmendezp/godesde0/ejercicios"
	"github.com/hmendezp/godesde0/funciones"
	g "github.com/hmendezp/godesde0/goroutines"
	"github.com/hmendezp/godesde0/iteraciones"
	"github.com/hmendezp/godesde0/mapas"
	"github.com/hmendezp/godesde0/modelos"
	"github.com/hmendezp/godesde0/teclado"
	"github.com/hmendezp/godesde0/users"
	"github.com/hmendezp/godesde0/variables"
)

func conVida(vida bool) string {
	if vida {
		return "vivo"
	} else {
		return "muerto"
	}
}

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
	// files.LeoArchivo()

	funciones.Calculos()
	// funciones.CallSimpleClosure()

	/* MiFibon := funciones.Fibonacci()
	for i := 1; i < 11; i++ {
		fmt.Printf("Fibonnacci(%d)=%d\n", i, MiFibon())
	}

	MiFactor := funciones.Factorial()
	for i := 0; i < 13; i++ {
		fmt.Printf("%d!=%d\n", i, MiFactor())
	} */
	funciones.Exponencia(2)

	// arreglos_slices.MostrarArreglos()
	arreglos_slices.MostrarSlices()
	mapas.MostrarMapas()

	users.AltaUsuario()

	Pedro := new(modelos.Hombre)
	Pedro.SetVida(true)
	e.HumanosRespirando(Pedro)
	fmt.Println("Este humano esta ", conVida(Pedro.EstaVivo()))

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
	// Maria.SetVida(false)
	fmt.Println("Este humano esta ", conVida(Maria.EstaVivo()))

	d.VemosDefer()
	// d.EjemploPanic()

	canal1 := make(chan bool)
	go g.DeletreaNombre("Elmer Homero", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aquí")
}
