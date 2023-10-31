package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresNumber() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número 1 : ")
	if scan.Scan() {
		numero1, err = strconv.Atoi(scan.Text())
		if err != nil {
			panic("el dato ingresado no puede ser convertido a entero!" + err.Error())
		}
	}
	fmt.Println("Ingrese número 2 : ")
	if scan.Scan() {
		numero2, err = strconv.Atoi(scan.Text())
		if err != nil {
			panic("el dato ingresado no puede ser convertido a entero!" + err.Error())
		}
	}
	fmt.Println("Ingrese texto de leyenda : ")
	if scan.Scan() {
		leyenda = scan.Text()
	}
	fmt.Println(leyenda, numero1*numero2)
}

func Nada() {
	fmt.Println("Nada personal!")
}
