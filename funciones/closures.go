package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func CallSimpleClosure() {
	tabladel := 3
	MiTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}

func Fibonacci() func() int {
	resant := 0
	resact := 0
	nuevores := 0
	return func() int {
		if resant == 0 && resact == 0 {
			resact = 1
			return 1
		}
		nuevores = resant + resact
		resant = resact
		resact = nuevores
		return nuevores
	}
}

func Factorial() func() int {
	multant := 0
	secuencia := 0
	return func() int {
		if multant == 0 {
			multant = 1
			return multant
		}
		secuencia += 1
		multant = multant * secuencia
		return multant
	}
}
