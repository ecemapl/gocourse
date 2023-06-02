package main

import "fmt"

// Función que devuelve el signo de su argumento
// (-1 si es negativo, +1 si es positivo, y 0 en cualquier otro caso)
func devuelveSigno(param int) string {
	if param > 0 {
		return "+1"
	} else if param < 0 {
		return "-1"
	} else {
		return "0"
	}
}

func main() {
	fmt.Println(devuelveSigno(10))
	fmt.Println(devuelveSigno(-1))
	fmt.Println(devuelveSigno(0))
}
