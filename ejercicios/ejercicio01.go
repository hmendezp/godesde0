package ejercicios

import "strconv"

func Excer01(param1 string) (int, string) {
	var valorRet string
	i, erroneo := strconv.Atoi(param1)
	if erroneo != nil {
		return 0, "Hubo un error en Atoi()!"
	}
	if i > 100 {
		valorRet = "Es mayor a 100"
	} else {
		valorRet = "Es menor a 100"
	}
	return i, valorRet
}
