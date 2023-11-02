package funciones

import "fmt"

func Calculos() {
	// var num1, num2 int
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}
	fmt.Println(suma(10, 35))
}
