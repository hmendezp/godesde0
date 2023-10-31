package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DespliegaTabla() {
	var numero int
	var err error
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero para generar tabla de multiplicar : ")
		if scan.Scan() {
			numero, err = strconv.Atoi(scan.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	MiTabla := Tabla(numero)
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, MiTabla())
	}
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func MostrarTablaDel(numero int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
