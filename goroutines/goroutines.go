package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func DeletreaNombre(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
	canal1 <- true
}
